func bulbSwitch(n int) int {
    //初始全部关闭，也就全部为0，那么考虑任意一个整数，都可以拆成两个数的乘积，就算是1也可以拆成1*1，比如4可以拆成1*4  2*2
    //那么拿4号灯来举例，该灯在1和4的时候被开关，两者抵消，回归到关闭状态，但由于2*2的原因，第二步进行切换后，没人给他还原，所以他永远保持亮着的状态
    //那也即是说只要找到<=n的完全平方数就好了
    //另一个点就是给n开方再取整就是n的完全平方数，我服了
   return int(math.Sqrt(float64(n)))
}
