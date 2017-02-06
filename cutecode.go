func createNewCC(code int) error {
        var (
		b bool = false
		err error
	        newCode int
	)

	/* todo 枚举，暂弃用
	regs := phalgo.Config.GetStringSlice("cutecodereg.regs")
	for newCode = code + 1;b==true;newCode++{
		newStr := strconv.Itoa(newCode)
		slice_b := false
             for _,v :=range regs{
		     if strings.Contains(newStr,v){
			     slice_b = true
			     break
		     }
	     }
		if slice_b == false{
			b = true
		}else{
			err = cutecodeinfo.New(newCode)
		}
	}
	*/
	for newCode = code + 1;b==false;newCode++{
		newStr := strconv.Itoa(newCode)

		topFour := common.SubString(newStr,0,4)
		if cuteCodeReg(topFour) == true{
			bottomFour := common.SubString(newStr,2,4)
			if cuteCodeReg(bottomFour) == true{
				b = true
			}
		}
		if b == false{
			err = cutecodeinfo.New(newCode)
		}else{
			break
		}
	}
	err = paramsinfo.UpdateValueIntByKey("cutecode",newCode)
	common.GetChannelByName("cutecode") <- newCode
	return err
}

func cuteCodeReg(number string) bool {
	 b := false                                                               //todo true为非靓号可用，false为靓号
	delta := int(number[1])-int(number[0])
	if delta == -1 || delta == 0 || delta == 1{                              //todo 正反顺，豹子
		for i:=2;i<4;i++{
			if  delta != int(number[i])-int(number[i-1]){
				b = true
				break
			}
		}
	}else{
		b = true
	}
     return b 

