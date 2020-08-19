# 点到直线的距离

## 点到直线的距离

> org.locationtech.jts.algorithm.distance

对于点到直线的距离，有以下几种情况

![](https://cloud.yuhang.ch/post-jts-p2ls.png)

对于线段AB，线段外一点P到线段的最短距离距离d

令

$$
\angle PAB=\theta
$$

设定标志有AP在AB上的投影与AB的长度之比为

$$
r=\frac{|AP|\cos \theta}{|AB|}=\frac{AP·AB}{|AP||AB|}·\frac{|AP|}{|AB|}=\frac{AP·AB}{|AB|^2}
$$

有以下几种情况：

$$
\begin{array}{c}r \leqslant 0, d=P A \\ r \geqslant 1, d=P B \\ 0<r<1, d=P L(P L \perp A B)\end{array}
$$

对于第三种情形，有

$$
|PL|=|AP|·\sin \theta=|AP|·\frac {AP\times AB}{|AP||AB|}=\frac{AP\times AB}{|AB|}
$$

