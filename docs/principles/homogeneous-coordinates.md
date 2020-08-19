# 齐次坐标系

## 引入

在计算机图形学终有一句名言：

> _“齐次坐标表示是计算机图形学的重要手段之一，它既能够用来明确区分向量和点，同时也更易用于进行仿射（线性）几何变换。”——_ F.S. Hill, JR。

简单来说，齐次坐标就是用$$N+1$$维来代表$$N$$维坐标.

> Homogeneous coordinates, introduced by August Ferdinand Möbius, make calculations of graphics and geometry possible in projective space.Homogeneous coordinates are a way of representing N-dimensional coordinates with N+1 numbers.

$$
\begin{aligned}(x, y, w) & \Leftrightarrow &\left(\frac{x}{w}, \frac{y}{w}\right) \\ \text { Homogeneous } & & \text { Cartesian } \end{aligned}
$$

## 为什么称为齐次坐标

在笛卡尔与齐次转换过程中，我们发现：

$$
\begin{array}{l}\text { Homogeneous } \quad \text { Cartesian } \\ \qquad \begin{aligned}(1,2,3) & \Rightarrow \quad\left(\frac{1}{3}, \frac{2}{3}\right) \\(2,4,6) & \Rightarrow \quad\left(\frac{2}{6}, \frac{4}{6}\right) \quad=\left(\frac{1}{3}, \frac{2}{3}\right) \\(4,8,12) & \Rightarrow \quad\left(\frac{4}{12}, \frac{8}{12}\right)=\left(\frac{1}{3}, \frac{2}{3}\right) \\ \vdots & \vdots \\(1 a, 2 a, 3 a) & \Rightarrow \quad\left(\frac{1 a}{3 a}, \frac{2 a}{3 a}\right)=\left(\frac{1}{3}, \frac{2}{3}\right) \end{aligned}\end{array}
$$

这些点都指向了欧几里得点$$(\frac{1}{3},\frac{2}{3})$$,无论怎样缩放变化，都代表欧几里得空间中的同一个点。因此这些点是“_homogeneous_”的，因为它们在欧几里得空间或笛卡尔空间中代表了同一个点。换句话说，齐次坐标是尺度不变的。

{% hint style="info" %}
词典中，“homogeneous”意思为同种族的，同质的。
{% endhint %}

## 点的齐次表示

我们只需要在现有的坐标中新增一个,$$w$$,变量就可以构造一个齐次坐标。这样，一个笛卡尔坐标系中的点$$(X,Y)$$

在齐次坐标系中就变成了$$(x,y,w)$$，并且齐次坐标还可以再投影到笛卡尔坐标系中：

$$
\begin{array}{l}X=\frac {x} {w} \\ Y=\frac {y} {w}\end{array}
$$

## 直线的齐次表示

一般的直线可以由以下方程表示：

$$
a x+b y+c=0
$$

平面上直线的齐次表示形式为：

$$
(x,y,w)
$$
$$
xX+yY+w=0
$$

对于任意非$$0$$的$$k$$:

$$
\begin{array}{l}(x,y,w) \\ k(x,y,w)\end{array}
$$

表示同一条直线。

## 点线空间关系的表示

点在直线上有：

$$
x·l = l·x=0
$$

两点确定的直线：

$$
x=l_{1} \times l_{2}
$$

两个直线的交点：

$$
l=x_{1} \times x_{2}
$$

## 参考

[Homogeneous Coordinates](http://www.songho.ca/math/homogeneous/homogeneous.html)

[关于齐次坐标的理解（经典）](https://blog.csdn.net/jeffasd/article/details/77944822)

[平面上点和直线的齐次表示](https://blog.csdn.net/qq_30622145/article/details/80837168)

