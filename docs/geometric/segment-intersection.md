---
title: 阅读JTS (数学重修) 之路
date: '2020-05-03T04:03:02.000Z'
draft: false
tags:
  - its
  - golang
categories:
  - notes
---

# 线段交点

### 求两线段交点

> org.locationtech.jts.algorithm.Intersection

对任意相交折线段，AB ，CD ，求其交点T：

![&#x5176;&#x4EA4;&#x70B9;](https://cloud.yuhang.ch/post-jts-inters-1.png)

### 找交点所在的矩形范围

对于直线

$$
l_1=AB，l_2=CD
$$

找到其对应的矩形范围

$$
l_1\in D_1,D_2\in l_2
$$

得到

$$
D = D_1\bigcap D_2
$$

进一步得到D的中心点M，将M作为原点建立坐标系，平移各点后得到

$$
A_m,B_m,C_m,D_m
$$

### 根据齐次坐标求交点坐标

高中的平面几何中，都在使用方程组的方法求解，相比之下，齐次坐标的形式更方便、简洁。

对齐次坐标有：

两点确定的直线为

$$
l=P_1\times P_2
$$

两直线交点为

$$
T=l_1\times l_2
$$

将坐标转换为齐次坐标有：

$$
A_m = (x_a,y_a,w_a)
$$

$$
B_m = (x_b,y_b,w_b)
$$

取 

$$
w_a = w_b = 1
$$

则

$$
l_1 = A_m\times B_m=(y_a-y_b,-(x_a-x_b),x_a\times y_b-x_b\times y_a)
$$

同理得：

$$
l_2 = C_m\times D_m=(y_c-y_d,-(x_c-x_d),x_c\times y_d-x_d\times y_c)
$$

有齐次坐标

$$
T=(x,y,w)=l_1\times l_2
$$

最终有Target T 坐标

$$
(X,Y) = (\frac{x}{w}+x_m,\frac{y}{w}+y_m)
$$

## 

