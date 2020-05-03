# 向量的计算

## 数量积

[数量积](https://zh.wikipedia.org/wiki/%E6%95%B0%E9%87%8F%E7%A7%AF)也叫点积，它是向量与向量的乘积，其结果为一个标量（非向量）。几何上，数量积可以定义如下：

设$$\vec{a}$$、$$\vec{b}$$ 两个任意向量，它们的夹角为$$\theta$$，则他们的数量积为：

$$
\vec{a} \cdot \vec{b}=|\vec{a}||\vec{b}| \cos \theta
$$

即 $$\vec{a}$$ 向量在$$\vec{b}$$ 向量方向上的投影长度（同方向为正反方向为负号），与$$\vec{b}$$向量长度的乘积。

## 向量积

[向量积](https://zh.wikipedia.org/wiki/向量积)也叫[叉积](https://zh.wikipedia.org/wiki/叉积)，[外积](https://zh.wikipedia.org/wiki/外积)，它也是向量与向量的乘积，不过需要注意的是，它的结果是个向量。它的几何意义是所得的向量与被乘向量所在平面垂直，方向由右手定则规定，大小是两个被乘向量张成的平行四边形的面积。所以向量积不满足交换律。

设有向量$$\vec{a}=a_{x} \vec{i}+a_{y} \vec{j}+a_{z} \vec{k}, \vec{b}=b_{x} \vec{i}+b_{y} \vec{j}+b_{z} \vec{k}$$

则其向量积的矩阵表达式可用下列符号表示：
$$
\vec{a} \times \vec{b}=\left|\begin{array}{ccc}\vec{i} & \vec{j} & \vec{k} \\ a_{x} & a_{y} & a_{z} \\ b_{x} & b_{y} & b_{z}\end{array}\right|
$$

### 坐标表示

向量$$\mathbf{u},\mathbf{v}$$可以定义为平行于基向量的三个正交元素之和：
$$
\begin{array}{l}\mathbf{u}=u_{1} \mathbf{i}+u_{2} \mathbf{j}+u_{3} \mathbf{k} \\ \mathbf{v}=v_{1} \mathbf{i}+v_{2} \mathbf{j}+v_{3} \mathbf{k}\end{array}
$$
$$\mathbf{s}=s_{1} \mathbf{i}+s_{2} \mathbf{j}+s_{3} \mathbf{k}=\mathbf{u} \times \mathbf{v}$$的三个标量元素可表示为
$$
\left(\begin{array}{l}s_{1} \\ s_{2} \\ s_{3}\end{array}\right)=\left(\begin{array}{l}u_{2} v_{3}-u_{3} v_{2} \\ u_{3} v_{1}-u_{1} v_{3} \\ u_{1} v_{2}-u_{2} v_{1}\end{array}\right)
$$

### 几何意义

如果以向量为边构成一个平行四边形，那么这两个向量叉积的模长与这个平行四边形的正面积相等。
$$
\|\mathbf{a} \times \mathbf{b}\|=\|\mathbf{a}\|\|\mathbf{b}\| \sin \theta
$$




