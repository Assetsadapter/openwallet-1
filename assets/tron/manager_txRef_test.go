/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package tron

import (
	"fmt"
	"testing"
)

var (
	pTxRaw    = "0a7d0a02d0762208239cf236e19b41cf40e887c8a7e12c5a66080112620a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412310a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518904e"
	pTxSigned = "0a7d0a02d0762208239cf236e19b41cf40e887c8a7e12c5a66080112620a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412310a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518904e124123218536461fff784bd32e206a4c1a7adc05b455aa37cd624724cb3a2826119d434317d121bda9b5352bf0aaa61326471c5167d14376a96a466317d041696c6801"
)

func TestCreateTransactionRef(t *testing.T) {

	var to_address, owner_address string
	var amount int64

	// predictTxRaw := "0a7e0a021031220816b0c1a29ce3387c40e099ad83e02c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a1541b6c1abf9fb31c9077dfb3c25469e6e943ffbfa7a12154199fee02e1ee01189bc41a68e9069b7919ef2ad8218c0843d"
	// to_address, owner_address, amount = "TQ1TiUzStbSLdEtACUDmzfMDpWUyo8cyCf", "TSdXzXKSQ3RQzQ5Ge8TiYfMQEjofSVQ8ax", uint64(1)
	to_address, owner_address, amount = TOADDRESS, OWNERADDRESS, AMOUNT

	if r, err := tw.CreateTransactionRef(to_address, owner_address, amount); err != nil {
		t.Errorf("TestCreateTransaction failed: %v\n", err)
	} else {
		t.Logf("TestCreateTransaction return: \n%+v\n", r)

		fmt.Println("")
		fmt.Println("APP: 0a7e0a02adcd220873bf2dfd044f459440e0d4d2f8e12c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518c0843d")
		fmt.Printf("Ref: %+v\n\n", r)
	}
}

func TestSignTransactoinRef(t *testing.T) {

	var txRaw string
	txRaw = "0a7e0a02a46f22084257d74387c4c4074090e9f1b1e62c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518c0843d"
	txRaw = "0a7e0a0222c32208d584f987611b830a40e1bfa8b2e62c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518c0843d"
	if r, err := tw.SignTransactionRef(txRaw, PRIVATEKEY); err != nil {
		t.Errorf("SignTransactionRef failed: %v\n", err)
	} else {
		t.Logf("SignTransactionRef return: \n\t%+v\n", r)
		debugPrintTx(r)
	}
}

func TestValidSignedTransactionRef(t *testing.T) {
	var txSignedRaw string
	txSignedRaw = TXSIGNED
	txSignedRaw = "0a7e0a02adcd220873bf2dfd044f459440e0d4d2f8e12c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518c0843d1241f5b8eaac6034f590b54d7d1e2fcd588c56573c113ec7e98aac4a393747ae290e55f1bc2e861cc9dde18ac48e9594054632f4a1da2491bf091c2fe813f4e373d201"
	txSignedRaw = "0a7e0a0269bb220876979998ddefc0d94088a5b192e62c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518e091431241e99894c32ec9b8d878c6624416aab59cc9bd0eeafb34d9e6d5a95b3472899a817dc87172dc2dec525cc99ae25247388a390e6827ee4fcaa9d29447dc24a8a3d100"

	if err := tw.ValidSignedTransactionRef(txSignedRaw); err != nil {
		t.Errorf("ValidSignedTransactionRef: %v\n", err)
	} else {
		t.Logf("GetTransactionSignRef return: \n\t%+v\n", "Success!")
	}
}

func TestSuiteTx(t *testing.T) {

	println("Start testsuit...\n")

	var (
		txRaw, txSignedRaw string
		err                error
	)

	txRaw, err = tw.CreateTransactionRef(TOADDRESS, OWNERADDRESS, AMOUNT)
	if err != nil {
		t.Errorf("TestCreateTransaction failed: %v\n", err)
	}
	println("txRaw: ", txRaw)
	println("------------------------------------------------------------------- Create Done! \n")

	// txRaw = "0a7e0a02adcd220873bf2dfd044f459440e0d4d2f8e12c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518c0843d"
	txSignedRaw, err = tw.SignTransactionRef(txRaw, PRIVATEKEY)
	if err != nil {
		t.Errorf("GetTransactionSignRef failed: %v\n", err)
	}
	println("txSignedRaw: ", txSignedRaw)
	println("------------------------------------------------------------------- Sign Done! \n")

	// txSignedRaw = "0a7e0a02adcd220873bf2dfd044f459440e0d4d2f8e12c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518c0843d1241f5b8eaac6034f590b54d7d1e2fcd588c56573c113ec7e98aac4a393747ae290e55f1bc2e861cc9dde18ac48e9594054632f4a1da2491bf091c2fe813f4e373d201"
	err = tw.ValidSignedTransactionRef(txSignedRaw)
	if err != nil {
		t.Errorf("ValidSignedTransactionRef: %v\n", err)
	} else {
		println("Success!")
	}
	println("------------------------------------------------------------------- Valid Done! \n")

	// txSignedRaw := "0a7e0a02b73b2208d971acfd3452661c40f0d68cfce12c5a67080112630a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412320a154199fee02e1ee01189bc41a68e9069b7919ef2ad82121541e11973395042ba3c0b52b4cdf4e15ea77818f27518c0843d1241166ae365792c1918de963cc4121e47417252da11d54783dbeb248f913240f27ea02b1b42f807c4ffb5d7ebecf687f5294400281021e6fefd0f38c50765f9c87200"
	err = tw.BroadcastTransaction(txSignedRaw)
	if err != nil {
		t.Errorf("ValidSignedTransactionRef: %v\n", err)
	} else {
		println("Success!")
	}
	println("------------------------------------------------------------------- Boradcast! \n")
}
