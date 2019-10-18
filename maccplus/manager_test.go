/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package maccplus

import (
	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/log"
	"github.com/shopspring/decimal"
	"path/filepath"
	"testing"
)

var (
	tw *WalletManager
)

func init() {

	tw = testNewWalletManager()
}

func testNewWalletManager() *WalletManager {
	wm := NewWalletManager()

	//读取配置
	absFile := filepath.Join("conf", "conf.ini")
	//log.Debug("absFile:", absFile)
	c, err := config.NewConfig("ini", absFile)
	if err != nil {
		return nil
	}
	wm.LoadAssetsConfig(c)
	//wm.ExplorerClient.Debug = false
	wm.WalletClient.Debug = true
	return wm
}

func TestWalletManager(t *testing.T) {

	t.Log("Symbol:", tw.Config.Symbol)
	t.Log("ServerAPI:", tw.Config.ServerAPI)
}

func TestGetCoreWalletinfo(t *testing.T) {
	tw.GetCoreWalletinfo()
}

func TestGetBlockChainInfo(t *testing.T) {
	b, err := tw.GetBlockChainInfo()
	if err != nil {
		t.Errorf("GetBlockChainInfo failed unexpected error: %v\n", err)
	} else {
		t.Logf("GetBlockChainInfo info: %v\n", b)
	}
}

func TestListUnspent(t *testing.T) {
	utxos, err := tw.ListUnspent(0, "MTeUvoBeLpaMatXE6P9GvTmW2pSFt6kNLP")
	if err != nil {
		t.Errorf("ListUnspent failed unexpected error: %v\n", err)
		return
	}
	totalBalance := decimal.Zero
	for _, u := range utxos {
		t.Logf("ListUnspent %s: %s = %s\n", u.Address, u.AccountID, u.Amount)
		amount, _ := decimal.NewFromString(u.Amount)
		totalBalance = totalBalance.Add(amount)
	}

	t.Logf("totalBalance: %s \n", totalBalance.String())
}

func TestEstimateFee(t *testing.T) {
	feeRate, _ := tw.EstimateFeeRate()
	t.Logf("EstimateFee feeRate = %s\n", feeRate.StringFixed(8))
	fees, _ := tw.EstimateFee(10, 2, feeRate)
	t.Logf("EstimateFee fees = %s\n", fees.StringFixed(8))
}

func TestGetNetworkInfo(t *testing.T) {
	tw.GetNetworkInfo()
}

func TestPrintConfig(t *testing.T) {
	tw.Config.PrintConfig()
}

func TestRestoreWallet(t *testing.T) {
	keyFile := "/myspace/workplace/go-workspace/projects/bin/data/btc/key/MacOS-W9JyC464XAZEJgdiAZxUXbPpsZZ2JeAujV.key"
	dbFile := "/myspace/workplace/go-workspace/projects/bin/data/btc/db/MacOS-W9JyC464XAZEJgdiAZxUXbPpsZZ2JeAujV.db"
	datFile := "/myspace/workplace/go-workspace/projects/bin/testdatfile/wallet.dat"
	tw.LoadConfig()
	err := tw.RestoreWallet(keyFile, dbFile, datFile, "1234qwer")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
	}

}

func TestWalletManager_ImportAddress(t *testing.T) {
	addr := "Ga2thK76EF4Y1q14RtmCfBZepC2GYBvaCy"
	err := tw.ImportAddress(addr, "")
	if err != nil {
		t.Errorf("RestoreWallet failed unexpected error: %v\n", err)
		return
	}
	log.Info("imported success")
}

func TestWalletManager_ListAddresses(t *testing.T) {
	addresses, err := tw.ListAddresses()
	if err != nil {
		t.Errorf("GetAddressesByAccount failed unexpected error: %v\n", err)
		return
	}

	for i, a := range addresses {
		t.Logf("ListAddresses address[%d] = %s\n", i, a)
	}
}

func TestWalletManager_GetInfo(t *testing.T) {
	tw.GetInfo()
}