func multiply(num1 string, num2 string) string {
    
    //这题我真不会，然后发现有个叫错位相乘的规则，就是两个数相乘=两个数的每一位数字的乘积，这我真不知道，小学乘法没学好╮(╯▽╰)╭
    //那知道这个规则就好办了，一位一位的乘，十位数字要累加后再取
    
    lens1 := len(num1)
    lens2 := len(num2)
    
    if lens1 == 0 || lens2 == 0 {
        return "0"
    }
    res := make([]byte, lens1+lens2)
    for k,_ := range res {
        res[k] = '0'
    }
    
    for i:=lens1-1; i >=0 ; i-- {
        for j:=lens2-1; j >= 0;j-- {
            n1 := num1[i] - '0'
            n2 := num2[j] - '0'
            sum := n1 * n2 + (res[i+j+1] - '0') //两个数相乘后要加上之前的十位，之前的十位就会变成现在的个位
            //然后现在的这个数要重新赋值十位个位，重点是现在的十位可能也有数了，那就要加上这个十位的数
            //我就很奇怪这个算法最后的结果怎么不会出现十位是两位数的，那样就炸了啊
            //这样，如果本次十位上有值的话，那么本次的计算一定不是最后一轮计算，
            //到了最后一轮计算的时候，那两个数的相乘，乘积+之前的十位也就是现在的个位，得出一个两位数，这个两位数填充
            //完个位，填充十位的时候，这个十位一定是空的，因为最后一轮的最后一组数字相乘，填充十位时这个十位一定是在
            //最前面的，所以此时即便 目前的十位+=之前这个位置的值，那么+的也只是个0而已，保证了最后的十位一定是10以内的数
            //也就是单个数
            res[i+j] =  ((res[i+j]-'0') + sum/10) + '0'
            res[i+j+1] = sum%10 + '0'
        }
    }
    
    //去掉前置0
    index := 0
    find := 0
    for i:=0; i < lens1+lens2; i++ {
        if res[i] != '0' {
            find = 1
            index = i
            break
        }
    }
    if index == 0 && find != 1 {
        return "0"
    }
    
    return string(res[index:])
    
    
}
