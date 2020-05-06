# 道格拉斯-普克算法

> org.locationtech.jts.simplify.DouglasPeuckerLineSimplifier

经典的线化简算法。1973年道格拉斯（David Douglas）和普克（Thomas Peucker）二人提出了该算法。其优点是具有平移和旋转不变性，给定曲线与阈值后，抽样结果一定。

计算步骤如下：

* 将目标折线首尾相连，找到与其连接线距离最远的点。
* 判断：若该距离小于所要求的阈值，则所有中间节点都被舍去。
* 否则，将该点标记为保留，并将区间一分为二，对两部分重复上述步骤。

{% hint style="info" %}
Golang 实现在 gots/algorithm/simplify/dp.go
{% endhint %}

### 参考

[道格拉斯-普克算法（Douglas–Peucker algorithm）](https://blog.csdn.net/baimafujinji/article/details/6475432)

