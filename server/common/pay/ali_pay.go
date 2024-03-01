package pay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
)

type QueryRes struct {
	PaySuccess     bool
	BuyerPayAmount string
}

func TradePagePay(title, orderNum, price string) (string, error) {
	privateKey := "MIIEogIBAAKCAQEAldCUT3Qg1sllMPSC/DlhdoZdAN96HVwfpBRxQnCHdyQTUyTm8z3wqSn0oNe2ifMR/YyAk1JXxm3/4d9XLnj/ODM4v/23kiNdW5kieCPEft/oyzNvMkeRZOmZpp4Ur8/4UhrbqOTXRaRzXpqynavatlzFg8xovUvvP/7WHGGbc06GNVWc8eDddGqmpcAzCXdOlmzRwzsqEQ+1unsXSPOuRJtzWMLvP0gjf0DT7e3tFpZjOo9Efrf4F3z/D8CiFnF7ePhp91FiaMHUMBR4BmwByufpWPORKYF12+yfN8R9CTuGjPyO95dvlVWUyvpItdYXe+G2VLADwRjjgPDDmksgtwIDAQABAoIBABNPz85uKc2GKIUm/7uDqgydRV/XbuZZ+bos6GyoFM5G2DbfiuXQFjW4MtCI/f7KVodYlTpoYugxRcrX/2T+M+gKskYNbpxn8qDcHJVgBvRT6K5I8wyILTXU5vmTZxdCt7/dtnMPDLDtgfGv/1ZT0N2xF8TllJrSLwUQ1IjUDDXQtYRsKJa86x/PjujGLAQ8WzVbImeec0odYCpbgu9kYS4Q8ozK/ELfKI+HdXvULKI4j4J/jhqNvt2NO6M97XqG6W/NJwT1HuioBts9NKs6tsZhBbhmLPoznVlahG1cxicEHtiGLU+S7JNLRzayhn9XjAu8NGAjmILngexfWdHgcAECgYEA+U6WQyw82IbZdn0Ymh+MUo+hHytZElgzc7Dyw+WLvK2jPFqENR6hNWki9i9EL18MoKg1BG2VmJpn1F4kvUrP85V0cFj3FU7D4Wuipmvd3hMiacL25AFtAv1Rfu+Ub9qqhT8kijuqIOjYV+cWWGQrHsNDitUsrCLGbR/pasV4mv8CgYEAmdY2Oqdpx++q7IMq2Jkt6RDJCchXlI2wT2DvrjWbzh8lxZtpHo5d2zeSE/K76Ky2zGsbr8K276GQhlDVkaFQ7QTFORHdF9XnLNmno8W5txQkKXtx2MgP2zQvjpayXoWHIWffmWiGotlXlgreNKBFxyrKxcpsStC+7BRiT5T/EkkCgYAKN7xU2HVN8ojrIEln168qmxRhcKdf5xmEvlXUzY2imAM+r/8HJlA5Hiv+pwVGY8DivMIfglZFQN7Mh2c+EtjfBNsW166YwveoP38VDwqy1Vntt73irmRHYRZ7+1m8E9w1Rdc7iyJ34tMrFx+FU+Wohxw3jnEcRPiY9FvesmU1BQKBgF3r1rgAtiiHPsefXW1YOBtvUA5U7oOX7ejIugyWVkWF5r26oHczXEIWr2zWabsol4I+cmzgaGftMBuRBpMe9hV/MBojueRvh/BuAleOxE/uUBWzdFJx8pxWRCK+BrCGvXW7wVwb4PdT0pGcmQVwRKnQinQvneYfhDEGUO9dfE1hAoGAILbNadyQOJYBi4M38Zz5NGFIKwVlfgrHAlRHNLuH0LDEvkW0LYg3c1cZ8d966LU32uwgqHfENcm+bQM61XEPdsPUpXdOMPDb4KZf/6EVx6GFJEvhYtuGwCrxsdJOO6EutKIdMfDA/GPl+v7XXfpPTJDJ+PqFakLGB1vJt2U8f28="
	client, err := alipay.NewClient("9021000129636518", privateKey, false)
	if err != nil {
		return "", err
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl("https://www.fmm.ink")

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", title)
	bm.Set("out_trade_no", orderNum)
	bm.Set("total_amount", price)

	//创建订单
	payUrl, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		return "", err
	}
	return payUrl, nil
}

func TradeQuery(orderNum string) (*QueryRes, error) {
	privateKey := "MIIEogIBAAKCAQEAldCUT3Qg1sllMPSC/DlhdoZdAN96HVwfpBRxQnCHdyQTUyTm8z3wqSn0oNe2ifMR/YyAk1JXxm3/4d9XLnj/ODM4v/23kiNdW5kieCPEft/oyzNvMkeRZOmZpp4Ur8/4UhrbqOTXRaRzXpqynavatlzFg8xovUvvP/7WHGGbc06GNVWc8eDddGqmpcAzCXdOlmzRwzsqEQ+1unsXSPOuRJtzWMLvP0gjf0DT7e3tFpZjOo9Efrf4F3z/D8CiFnF7ePhp91FiaMHUMBR4BmwByufpWPORKYF12+yfN8R9CTuGjPyO95dvlVWUyvpItdYXe+G2VLADwRjjgPDDmksgtwIDAQABAoIBABNPz85uKc2GKIUm/7uDqgydRV/XbuZZ+bos6GyoFM5G2DbfiuXQFjW4MtCI/f7KVodYlTpoYugxRcrX/2T+M+gKskYNbpxn8qDcHJVgBvRT6K5I8wyILTXU5vmTZxdCt7/dtnMPDLDtgfGv/1ZT0N2xF8TllJrSLwUQ1IjUDDXQtYRsKJa86x/PjujGLAQ8WzVbImeec0odYCpbgu9kYS4Q8ozK/ELfKI+HdXvULKI4j4J/jhqNvt2NO6M97XqG6W/NJwT1HuioBts9NKs6tsZhBbhmLPoznVlahG1cxicEHtiGLU+S7JNLRzayhn9XjAu8NGAjmILngexfWdHgcAECgYEA+U6WQyw82IbZdn0Ymh+MUo+hHytZElgzc7Dyw+WLvK2jPFqENR6hNWki9i9EL18MoKg1BG2VmJpn1F4kvUrP85V0cFj3FU7D4Wuipmvd3hMiacL25AFtAv1Rfu+Ub9qqhT8kijuqIOjYV+cWWGQrHsNDitUsrCLGbR/pasV4mv8CgYEAmdY2Oqdpx++q7IMq2Jkt6RDJCchXlI2wT2DvrjWbzh8lxZtpHo5d2zeSE/K76Ky2zGsbr8K276GQhlDVkaFQ7QTFORHdF9XnLNmno8W5txQkKXtx2MgP2zQvjpayXoWHIWffmWiGotlXlgreNKBFxyrKxcpsStC+7BRiT5T/EkkCgYAKN7xU2HVN8ojrIEln168qmxRhcKdf5xmEvlXUzY2imAM+r/8HJlA5Hiv+pwVGY8DivMIfglZFQN7Mh2c+EtjfBNsW166YwveoP38VDwqy1Vntt73irmRHYRZ7+1m8E9w1Rdc7iyJ34tMrFx+FU+Wohxw3jnEcRPiY9FvesmU1BQKBgF3r1rgAtiiHPsefXW1YOBtvUA5U7oOX7ejIugyWVkWF5r26oHczXEIWr2zWabsol4I+cmzgaGftMBuRBpMe9hV/MBojueRvh/BuAleOxE/uUBWzdFJx8pxWRCK+BrCGvXW7wVwb4PdT0pGcmQVwRKnQinQvneYfhDEGUO9dfE1hAoGAILbNadyQOJYBi4M38Zz5NGFIKwVlfgrHAlRHNLuH0LDEvkW0LYg3c1cZ8d966LU32uwgqHfENcm+bQM61XEPdsPUpXdOMPDb4KZf/6EVx6GFJEvhYtuGwCrxsdJOO6EutKIdMfDA/GPl+v7XXfpPTJDJ+PqFakLGB1vJt2U8f28="
	client, err := alipay.NewClient("9021000129636518", privateKey, false)
	if err != nil {
		return nil, err
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2)

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderNum)

	//查询订单
	aliRsp, err := client.TradeQuery(context.Background(), bm)
	if err != nil {
		return nil, err
	}
	var res QueryRes
	if aliRsp.Response.BuyerPayAmount == aliRsp.Response.TotalAmount {
		res.PaySuccess = true
		res.BuyerPayAmount = aliRsp.Response.BuyerPayAmount
	}
	return &res, nil
}
