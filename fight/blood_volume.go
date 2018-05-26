package fight

import "sync"

type basicState struct {
	PH     int           `json:"ph"`     //生命值
	Energy int           `json:"energy"` //能量值
	lock   *sync.RWMutex `json:"-"`
}
func newBasicState(ph,energy int) *basicState {
	return &basicState{ph, energy, new(sync.RWMutex)}
}

func (f *basicState) SetPH(redPH int) int {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.PH += f.PH
	return f.PH

}

func (f *basicState) GetPH() int {
	f.lock.RLock()
	defer f.lock.RLocker()
	return f.PH
}

func (f *basicState) SetEnergy(redEnergy int) int {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.Energy += f.Energy
	return f.Energy

}

func (f *basicState) GetEnergy() int {
	f.lock.RLock()
	defer f.lock.RLocker()
	return f.Energy
}
