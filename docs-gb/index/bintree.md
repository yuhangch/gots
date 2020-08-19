# 二叉树

首先从测试代码入手，看看其中的二叉树是怎么工作的。

首先定义了三个常量：测试的要素个数和范围的最大最小值。

```java
static final int NUM_ITEMS = 20000;
static final double MIN_EXTENT = -1000.0;
static final double MAX_EXTENT = 1000.0;
```

两个变量，分别是二叉树对象和一个`interval`的列表

```java
 IntervalList intervalList = new IntervalList();
 Bintree btree = new Bintree();
```

接着开始初始化网格。计算、获取以下值：

* 网格长度 = NUM\_ITEMS 的平方根 + 1
* extent的范围 = MAX\_EXTENT - MIN\_EXTENT
* 网格递增值 = extent的范围 / 网格长度

依据以上值开始循环填充树🌲。在每个循环中执行：

* 获取起始interval的最小值，最小值 = MIN\_EXTENT + 网格增长值 \* i
* 新建interval
* 将interval插入树中
* 将interval加入之前新建的interval列表

插入树中的逻辑如下：

Bintree中的insert为

```java
insert(Interval itemInterval, Object item)
```

其步骤为

* 收集interval的状态：将该interval的width与Bintree的min\_extent字段比较，若width大于0且小于min\_extent，则更新mi n\_extent字段更新为width的值。
* 确认范围是否有效：分析该interval的min与max，若两者不相等，认为有效，返回该interval，否则重新计算两个值。新的min与max为原始min+-（min\_extent/2\)
* 然后执行Bintree的root对象的插入，函数与Bintree中参数相同，如下：

  ```java
  insert(Interval itemInterval, Object item)
  ```

root对象执行插入操作步骤如下：

* 获取该interval在子节点中的索引。

  ```java
  int getSubnodeIndex(Interval interval, double centre)
  ```

  若interval的min大于centre，为右节点，返回1

  若interval的max小于centre，为左节点，返回0

  否则返回-1

* 若返回值为-1，则将item新加进root对象的items中，插入结束
* 否则，根据索引获取该子节点
* 若该子节点不存在或子节点的interval未包含该interval，则需要新建一个子节点。

  ```java
  Node createExpanded(Node node, Interval addInterval)
  ```

  当子节点不存在时：

  * 获取interval的值拷贝
  * 根据interval新建node并返回

  当子节点存在，需要对其expand：

  * 获取interval值拷贝，执行expand

    ```java
    expandToInclude(Interval interval)
    ```

    该方法步骤如下：

    若子节点max比该interval的max大，更新该interval的max为子节点的max

    若子节点的min比该interval的min小，更新该interval的min为子节点的min

    至此，完成该interval的expand

  * 根据interval新建node
  * 将子节点作为该node的子节点插入，返回该node。

* 将新生成的子节点替换原来的子节点。
* 将内容加入子节点

  ```java
  insertContained(Node tree, Interval itemInterval, Object item)
  ```

  其中tree为子节点。

  * 断言子节点的interval是否包含该interval，若否，断言失败。
  * 若该interval的width为0值或很小可认为是零值则逐层进行查找，直到返回拥有最小interval的节点。
  * 否则逐层查找节点，若不存在则新建该节点。
  * 将item添加到该节点。

