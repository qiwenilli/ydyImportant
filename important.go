package ydyImportant

/**
* qiwen<34214399@qq.com> 敏感数据过滤
*/
type Important struct {}

func (*Important) Mobile(str string) string{
    if(len(str)==11){
        return str[0:3]+"****"+str[7:]
    }else{
        return str
    }
}

func (*Important) BankCard(str string) string{
    if(len(str)==16||len(str)==19){
        return str[0:8]+"*******"+str[15:]
    }else{
        return str
    }
}

func (*Important) IdCard(str string) string{
    if(len(str)==15||len(str)==18){
        return str[0:4]+"**********"+str[14:]
    }else{
        return str
    }
}
