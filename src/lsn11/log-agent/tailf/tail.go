package tailf

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"lsn11/log-agent/conf"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type TailInfo struct {
	tail     *tail.Tail
	conf     conf.CollectConf
	exitChan chan int
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailInfoMgr struct {
	//tailInfoList []*TailInfo
	m         sync.Map
	msgChan   chan *TextMsg
	keyType   reflect.Type
	valueType reflect.Type
}

var (
	tailInfoMgr *TailInfoMgr
)

func NewTailInfoMgr(keyType, valueType reflect.Type) (tailfMgr *TailInfoMgr, err error) {
	if keyType == nil {
		err = errors.New("nil key type")
		return
	}

	if !keyType.Comparable() {
		err = errors.New("key incomparable")
		return
	}

	if valueType == nil {
		err = errors.New("nil value type")
		return
	}

	tailfMgr = &TailInfoMgr{
		keyType:   keyType,
		valueType: valueType,
		msgChan:   make(chan *TextMsg, conf.AppConfig.ChanSize),
	}
	return
}

func (p *TailInfoMgr) addTail(tail *TailInfo) (err error) {
	if reflect.TypeOf(tail.conf.LogPath) != p.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(tail.conf.LogPath)))
	}
	if reflect.TypeOf(tail) != p.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(tail)))
	}

	p.m.Store(tail.conf.LogPath, tail)
	err = nil
	return
}

func (p *TailInfoMgr) delTail(logPath string) (err error) {
	if reflect.TypeOf(logPath) != p.keyType {
		err = errors.New("error key type")
		return
	}

	p.m.Delete(logPath)
	return
}

func (p *TailInfoMgr) rangeTail(f func(key, value interface{}) bool) {
	p.m.Range(f)
}

func (p *TailInfoMgr) loadTail(logPath string) (tail *TailInfo, err error) {
	if reflect.TypeOf(logPath) != p.keyType {
		err = errors.New("error key type")
		return
	}

	val, ok := p.m.Load(logPath)
	if !ok {
		err = errors.New("load error")
		return
	}
	tail = val.(*TailInfo)
	return
}

func createNewTask(conf *conf.CollectConf) (err error) {
	tails, err := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		logs.Error("createNewTask failed err ", err)
		return
	}

	tailInfo := &TailInfo{
		conf:     *conf,
		tail:     tails,
		exitChan: make(chan int, 1),
	}

	tailInfoMgr.addTail(tailInfo)

	go readFromTail(tailInfo)
	return
}

func InitTail(collectConf []conf.CollectConf) (err error) {

	TailInfoMgrG, err := NewTailInfoMgr(reflect.TypeOf("string"), reflect.TypeOf(&TailInfo{}))
	if err != nil {
		panic(errors.New("init tailf mgr failed"))
	}
	tailInfoMgr = TailInfoMgrG

	if len(collectConf) == 0 {
		err = errors.New("invalid collectConf args")
		return
	}

	for _, val := range collectConf {
		createNewTask(&val)
	}
	return
}

func readFromTail(tailInfo *TailInfo) {
	for true {
		select {
		case line, ok := <-tailInfo.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, fileName %s\n", tailInfo.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}

			textMsg := &TextMsg{
				Msg:   line.Text,
				Topic: tailInfo.conf.Topic,
			}
			//logs.Debug("log line ", line)
			tailInfoMgr.msgChan <- textMsg
		case <-tailInfo.exitChan:
			logs.Warn("tailinfo will exited conf %v", tailInfo.conf)
			return
		}
	}
}

func GetLineMsg() (msg *TextMsg) {
	msg = <-tailInfoMgr.msgChan
	return
}

func UpdateConfig(conf []conf.CollectConf) (err error) {

	for _, oneConf := range conf {
		//查询更新后新的path是否存在goruntine在运行
		if _, ok := tailInfoMgr.loadTail(oneConf.LogPath); ok != nil {
			createNewTask(&oneConf)
		}
	}

	mm := tailInfoMgr.m
	tailInfoMgr.m.Range(func(key, value interface{}) bool {
		status := false
		for _, oneConf := range conf {
			if oneConf.LogPath == key.(string) {
				status = true
				break
			}
		}

		if status == false {
			val, ok := mm.Load(key.(string))
			if !ok {
				logs.Error("key has been delete")
			} else {
				val.(*TailInfo).exitChan <- 1
				mm.Delete(key.(string))
			}
		}

		return true
	})
	return
}
