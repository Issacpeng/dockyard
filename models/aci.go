package models

import (
//	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/redis.v3"

	"github.com/containerops/wrench/db"
)

type Aci struct {
	AciId         string   `json:"aciId"`         //
	Imagename     string   `json:"imagename"`     //
	Manifest      string   `json:"manifest"`      //
	SignPath      string   `json:"signPath"`      // 
	AciPath       string   `json:"aciPath"`       // 
	CreatedTime   int64    `json:"createTime"`    //
	UpdatedTime   int64    `json:"updatedTime"`   //
}


func (a *Aci) Has(aci string) (bool, string, error) {
	fmt.Printf("######### Aci Has : %v #########\r\n", aci)
	if key := db.Key("aci", aci); len(key) <= 0 {
		return false, "", fmt.Errorf("Invalid aci key")
	} else {
		if err := db.Get(a, key); err != nil {
			if err == redis.Nil {
				return false, "", nil
			} else {
				return false, "", err
			}
		}

		return true, key, nil
	}
}

func (a *Aci) Save() error {
	key := db.Key("aci", a.AciId)
	fmt.Printf("######### Aci Save key: %v #########\r\n", key)
	if err := db.Save(a, key); err != nil {
		return err
	}

	return nil
}

func (a *Aci) PutManifest(aciId, manifest string, size int64) error {
	if has, _, err := a.Has(aciId); err != nil {
		return err
	} else if has == false {
		a.AciId = aciId
		a.Manifest = manifest
		a.CreatedTime = time.Now().UnixNano() / int64(time.Millisecond)

		if err = a.Save(); err != nil {
			return err
		}
	} else {
	    a.Manifest, a.UpdatedTime = manifest, time.Now().UnixNano()/int64(time.Millisecond)

		if err := a.Save(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Aci) PutSignPath(aciId string, path string) error {
	if has, _, err := a.Has(aciId); err != nil {
		return err
	} else if has == false {
		return fmt.Errorf("Aci not found")
	} else {
		a.SignPath, a.UpdatedTime = path, time.Now().UnixNano()/int64(time.Millisecond)

		if err := a.Save(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Aci) PutAciPath(aciId string, path string) error {
	if has, _, err := a.Has(aciId); err != nil {
		return err
	} else if has == false {
		return fmt.Errorf("Aci not found")
	} else {
		a.AciPath, a.UpdatedTime = path, time.Now().UnixNano()/int64(time.Millisecond)

		if err := a.Save(); err != nil {
			return err
		}
	}

	return nil
}