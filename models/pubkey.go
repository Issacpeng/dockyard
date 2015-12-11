package models

import (
//	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"

	"gopkg.in/redis.v3"

	"github.com/containerops/wrench/db"
)

type Pubkey struct{
	PubkeyTag    string
	Pubkeyfile   []byte
	CreatedTime  int64
	UpdatedTime  int64
}

func (p *Pubkey) Has(pubkeyTag string) (bool, string, error) {
	fmt.Printf("######### Pubkey pubkeyTag : %v #########\r\n", pubkeyTag)
	if key := db.Key("pubkey", pubkeyTag); len(key) <= 0 {
		return false, "", fmt.Errorf("Invalid pubkeyTag")
	} else {
		if err := db.Get(p, key); err != nil {
			if err == redis.Nil {
				return false, "", nil
			} else {
				return false, "", err
			}
		}

		return true, key, nil
	}
}

func (p *Pubkey) Save() error {
	key := db.Key("pubkey", p.PubkeyTag)
	fmt.Printf("######### pubkey Save key: %v #########\\r\n", key)
	if err := db.Save(p, key); err != nil {
		return err
	}

	if _, err := db.Client.HSet(db.GLOBAL_ACI_INDEX, p.PubkeyTag, key).Result(); err != nil {
		return err
	}

	return nil
}

func (p *Pubkey)UploadPubkey(servername string, pubkeypath string) error{
	if pubkeypath == "" {
//		log.Error("[ACI PutPubkey] File path is invalid: %v", err.Error())

		return fmt.Errorf("Aci PutPubkey File path is invalid")
	}

	file, err := ioutil.ReadFile(pubkeypath)
	if err != nil {
//		log.Error("[ACI PutPubkey] Read PutPubkey file failed: %v", err.Error())

		return fmt.Errorf("Aci  Read PutPubkey file failed")
	}
    
	//Upload pubkey
	if has, _, err := p.Has(servername); err != nil {
		return err
	} else if has == false {
		p.PubkeyTag = servername
		p.Pubkeyfile = []byte(file) 
		p.CreatedTime = time.Now().UnixNano() / int64(time.Millisecond)

		if err = p.Save(); err != nil {
			return err
		}
	} else {
		p.PubkeyTag, p.Pubkeyfile, p.UpdatedTime  = servername, file, time.Now().UnixNano() / int64(time.Millisecond)

		if err := p.Save(); err != nil {
			return err
		}
	}

	return nil
}

func (p *Pubkey)GetPubkey(servername string) ([]byte, error) {
	if has, _, err := p.Has(servername); err != nil {
		return []byte(""), err
	} else if has == false {
		return []byte(""), fmt.Errorf("Pubkey not found")
	} else {
		return p.Pubkeyfile, nil
	}
}
