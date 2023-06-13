---
title: Games101-计算机图形学
created: 2023-06-10 13:48:15
---

## 计算机图形学

![](https://www.zhihu.com/equation?tex=1+2)

### 应用
- 电影、动画特效、设计（CAD）、CG（Computer Graphics）
- 可视化、虚拟现实、仿真/模拟
- 图形界面、字体...

### 挑战
真实世界的理解、计算方法、显示方法

### 内容
数学理论、着色、形体、仿真动画，不包括OpenGL、DirectX、Vulcan等（图形学api）

- 光栅化：广泛应用在实时图形学(>30fps)
- 几何（曲线和曲面）
- 光线追踪：能达到较好的效果，但较慢
- 动画/模拟

> 图形学（模型渲染图片）!=计算机视觉（图片提取信息：分析、猜测、理解、推理）

### 作业
不强制要求，意味着听听就行。每周具体作业代码不超过 20 行。最后有一个大作业。

## 线性代数
坐标、平移、旋转 ==> 矩阵计算

### 向量
![](https://www.zhihu.com/equation?tex=\bbox[white]{\vec{AB}=\vec{B}-\vec{A}})

- 向量加法（三角形法则）
![](https://www.zhihu.com/equation?tex=\bbox[white]{\vec{AB}=\vec{AO}%2B\vec{OB}})

- 向量点乘（Dot）: 表示方向性
![](https://www.zhihu.com/equation?tex=\bbox[white]{\vec{a}\cdot\vec{b}=|\vec{a}||\vec{b}|\cdot\cos(\theta)})

- 投影
b在a上的投影：![](https://www.zhihu.com/equation?tex=\bbox[white]{k=|\vec{b}|\cos\theta})

- 叉乘（方向根据右手定则）：已知两个坐标系得到另一个坐标系
![](https://www.zhihu.com/equation?tex=\bbox[white]{|\vec{a}\times\vec{b}|=|\vec{a}||\vec{b}|\sin\phi})


### 矩阵
- 矩阵乘法（矩阵X矩阵、矩阵X向量）
- 结合律、分配律、无交换律
- 单位矩阵
- 逆矩阵

## 变换
变换即是对目标的每一个点进行变换，以得到整体的改变。复杂变换可由简单变换得到，和变换的顺序有关。

### 二维变换

#### 缩放变换
![](https://www.zhihu.com/equation?tex=\bbox[white]{x^{'}=sx~~or~~y^{'}=sy})

#### 反转
![](https://www.zhihu.com/equation?tex=\bbox[white]{x^{'}=-x~~or~~y^{'}=-y})


$$
\begin{bmatrix}
x^{'} \\ y^{'}
\end{bmatrix}
=
\begin{bmatrix}
-1 & 0 \\ 0 & 1 
\end{bmatrix}
\begin{bmatrix}
x \\ y
\end{bmatrix}
$$

#### 切变

$$\begin{bmatrix}
x^{'} \\ y^{'}
\end{bmatrix}
=
\begin{bmatrix}
1 & a \\ 0 & 1 
\end{bmatrix}
\begin{bmatrix}
x \\ y
\end{bmatrix}
$$

#### 旋转

$$R_\theta=\begin{bmatrix}
\cos\theta & -\sin\theta\\
\sin\theta & \cos\theta 
\end{bmatrix}
$$

#### 线性变换

$$\begin{bmatrix}
x^{'}\\ y^{'}
\end{bmatrix}
=
\begin{bmatrix}
a&b\\ c&d
\end{bmatrix}
\begin{bmatrix}
x\\ y
\end{bmatrix}$$

即
$$x^{'}=ax+by$$

$$y^{'}=cx+dy$$

即
$$x^{'}=Mx$$

称为线性变换。以上均属于线性变换。

#### 平移
平移是一种特殊的二维变换，它不属于线性变换，因为：
$$
\begin{bmatrix}
x^{'}\\ y^{'}
\end{bmatrix}=
\begin{bmatrix}
a & b \\ c&d
\end{bmatrix}
\begin{bmatrix}
x \\ y
\end{bmatrix}
+
\begin{bmatrix}
t_x \\ t_y
\end{bmatrix}$$

为了解决其特殊性，引入齐次坐标。

#### 引入齐次坐标
为二维的点或向量增加一个维度，得到仿射变换的通式。

$$\begin{bmatrix}
x^{'}\\ y^{'} \\ w^{'}
\end{bmatrix}
=
\begin{bmatrix}
1&0&t_x\\ 0&1&t_y\\ 0& 0 &1
\end{bmatrix}
\begin{bmatrix}
x\\ y\\ 1
\end{bmatrix}
=\begin{bmatrix}
x+t_x \\ y+t_y\\ 1
\end{bmatrix}$$

w 为 0 或 1，以满足平移变换后的向量或点的不变性。其中 0 表示向量，1 表示点。

因此：
- 点-点=向量
- 向量+-向量=向量
- 点+-向量=点
- 点+点=2？=>点，若w!=0，w=w/w，x=x/w，其他同理

代价即是引入了一个额外的坐标，但是可省略仿射变换，只保留左上角矩阵，因此代价不高。

另外，在三维空间中，用四个参数描述一个三维齐次坐标系下的点。

#### 变换的组合性

$$R=A_1*A_2*...*S$$

复杂变换可由简单变换组合而成，其顺序相关，同时变换矩阵维数相同（3\*3）

### 三维变换
同二维变换，用四个坐标描述一个三维坐标中的点：一个四维矩阵

#### 旋转

有旋转矩阵：

$$
R_{x}(\alpha)=
\begin{bmatrix}
1&0&0&0 \\
0&\cos \alpha & -\sin \alpha &0 \\
0 &\sin \alpha &\cos \alpha &0 \\
0&0&0&1
\end{bmatrix}
$$

$$
R_{y}(\alpha)=
\begin{bmatrix}
\cos \alpha&0&\sin \alpha&0 \\
0&1 & 0 &0 \\
- \sin \alpha&0 &\cos \alpha &0 \\
0&0&0&1
\end{bmatrix}
$$

$$
R_{z}(\alpha)=
\begin{bmatrix}
\cos \alpha&-\sin \alpha&0 &0\\
\sin \alpha&\cos \alpha & 0 &0 \\
0 &0 &0&0 \\
0&0&0&1
\end{bmatrix}
$$

三维旋转可由三个独立二维旋转行为描述，用欧拉角描述，分为横滚角、俯仰角、航向角

$$R_{xyz}(\alpha,\beta,\gamma)=R_x(\alpha)R_y(\beta)R_z(\gamma)$$

> 罗德里格斯旋转公式：
> $$R = I + \sin\theta * A + (1-\cos\theta) * A^2$$
> 可计算出任意轴和任意角度的旋转变换矩阵

### MVP变换
M - Model 变换: 将物体从对象空间转换到世界空间。它包含物体的缩放、旋转和位置变换。

V - View 变换: 将世界空间转换到视图空间。它决定摄像机的位置和方向。

P - Projection 变换: 将视图空间转换到裁剪空间。它实现透视效果, 包含视锥体和视口的定义。

这三个变换的综合结果是将三维物体最终投影到二维屏幕上的效果。它们的作用顺序是: 物体空间 → 世界空间 → 视图空间 → 裁剪空间 → 屏幕空间。

### 投影
投影分为正交投影和透视投影，正交投影能够更好地反映比例，透视投影更加贴近真实。

方法是定义近和远两个平面，从一个平面投影向另一个平面。正交投影和透视投影的区别在于投影线是否平行，也就是投影面是否等大。

正交的计算方法形同平移，投影的坐标计算依赖于相似三角形。

$$y^{'}=\frac{n}{z}y$$

$$x^{'}=\frac{n}{z}x$$

n 和 z 分别是远平面和近平面到延长线和水平面交点的距离。

一个很好的资料是：[（超详细！）计算机图形学 入门篇 2. 视图变换（Viewing Transformations）](https://zhuanlan.zhihu.com/p/448547679)

#### 正交投影
正交投影的工作是将空间的长方体移动到原点并压缩成一个正方体。

有

$$
M_{ortho}=\begin{bmatrix}
\frac{2}{r-l}&0&0&0 \\
0 & \frac{2}{t-b} &0&0\\ 
0&0& \frac{2}{n-f} &0 \\
0&0&0&1
\end{bmatrix}
\begin{bmatrix}
1&0&0&- \frac{r+l}{2} \\
0&1&0&- \frac{t+b}{2} \\
0&0&1&- \frac{n+f}{2}\\
0&0&0&1 \\
\end{bmatrix}
$$

其中，(l,r)(b,t)(f,n)定义一个长方体，代表左右、上下、前后。

#### 透视投影
透视投影的工作是将远平面压缩成近平面等大，即将一个梯台变换为一个长方体。之后再进行正交投影就称为透视变换。

有

$$
M_{persp}=\begin{bmatrix}
n&0&0&0 \\
0&n&0&0 \\
0&0&n+f&-fn \\
0&0&1&0
\end{bmatrix}
$$

#### 总变换矩阵
先后进行透视投影和正交投影，得到总变换矩阵：

$$
M_{per}=M_{ortho}M_{persp}=\begin{bmatrix}
\frac{2n}{r-l}&0& \frac{l+r}{l-r}&0 \\
0& \frac{2n}{t-b} & \frac{b+t}{b-t}&0 \\
0&0& \frac{f+n}{n-f} & \frac{2fn}{f-n} \\
0&0&1&0 \\
\end{bmatrix}
$$

#### 另一种表述
有时我们偏向使用eye_fovy（垂向视角）和aspect_ratio（宽高比进行表述）。其中，

$$\frac{t}{n}=\tan{\frac{fovy}{2}}$$

$$aspect=\frac{t}{r}$$


### PA0
P点坐标（2，1），逆时针旋转45度，再平移（2，1），利用齐次坐标计算变换后的坐标。

```cpp
#include <cmath>
#include <eigen3/Eigen/Core>
#include <eigen3/Eigen/Dense>
#include <iostream>
#include <math.h>
#include <numbers>
#include <ostream>

using namespace std::numbers;

int main() {
  // 定义齐次坐标3维向量
  auto p = Eigen::Vector3f{2, 1, 1};
  //定义变换矩阵(3D)
  Eigen::Matrix3f r;
  r << cos(pi / 4), -sin(pi / 4), 2, sin(pi / 4), cos(pi / 4), 1, 0, 0, 1;
  std::cout << "Print R: " << std::endl;
  std::cout << r << std::endl;
  //计算变换后坐标
  p = r * p;
  //抛弃第3维，保留二维结果
  std::cout << "result print here" << std::endl;
  std::cout << "(" << p.x() << "," << p.y() << ")" << std::endl;
}
```

输出如下：
```sh
Print R: 
 0.707107 -0.707107         2
 0.707107  0.707107         1
        0         0         1
result print here
(2.70711,3.12132)
```

PA0 很容易，这里借 PA0 大致说一下环境配置问题：
1. 需要 eigen 和 opencv 两个库，推荐使用系统的包管理器下载: `paru -S eigen opencv`
2. opencv 可能由于路径问题（多出一个 opencv4 目录）无法找到库，建立一个软链接：`sudo ln -s /usr/include/opencv4/opencv2 /usr/include/`
3. 推荐使用 xmake 管理项目（当然你用 vs 或 cmake 也行），需要在 xmake. lua 中添加依赖 `set_requires("opencv")` 以及 `set_packages("opencv")`
4. 然后使用xmake编译和运行

### PA1
填写一个旋转矩阵和一个透视投影矩阵。给定三维下三个点 v0 (2.0, 0.0, −2.0), v1 (0.0, 2.0, −2.0), v2 (−2.0, 0.0, −2.0), 你需要将这三个点的坐标变换为屏幕坐标并在屏幕上绘制出对应的线框三角形。

> PS: 大概是我太弱，PA1 做得异常艰难，好不容易才弄懂。重要的大概就是几个变换矩阵，知道怎么推出来直接拿来用就好。再次鸣谢 [keanu大佬的笔记](https://zhuanlan.zhihu.com/p/448547679)！

模型变换代码如下：
```cpp
Eigen::Matrix4f get_model_matrix(float rotation_angle) {
  // 创建一个单位阵
  Eigen::Matrix4f model = Eigen::Matrix4f::Identity();
  Eigen::Matrix4f rotation;

  // 角度制转弧度制
  rotation_angle = rotation_angle * std::numbers::pi / 180;
  // 返回一个旋转矩阵，同齐次坐标的矩阵相乘
  rotation << cos(rotation_angle), -sin(rotation_angle), 0, 0,
      sin(rotation_angle), cos(rotation_angle), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1;

  model = rotation * model;
  return model;
}
```

投影变换代码如下：
```cpp
// 给出的参数包括眼角和宽高比，近平面和远平面
Eigen::Matrix4f get_projection_matrix(float eye_fov, float aspect_ratio,
                                      float zNear, float zFar) {

  Eigen::Matrix4f projection = Eigen::Matrix4f::Identity();

  // 角度制转弧度制
  eye_fov = eye_fov * std::numbers::pi / 180;

  // 倒三角问题
  // 注意，此处的zNear和zFar传入的均是绝对值，表示相对屏幕的距离。需要转化为坐标
  // 由于看向的是-z方向，最后的变换矩阵w的1也需要改为-1

  float n = -zNear;
  float f = -zFar;
  float t = n * tan(eye_fov / 2);
  float b = -t;
  float r = t / aspect_ratio;
  float l = -r;

  Eigen::Matrix4f per;
  per << 2 * n / (r - l), 0, (l + r) / (l - r), 0, 0, 2 * n / (t - b),
      (b + t) / (b - t), 0, 0, 0, (f + n) / (n - f), 2 * f * n / (f - n), 0, 0,
      -1, 0;
  projection = per * projection;
  return projection;
}
```

## 光栅化
### 一些定义
#### 屏幕
二维数组，表示屏幕大小（像素多少），称为分辨率；典型的光栅（Raster，德语的 Screen）成像设备。

#### 像素（Pixel）
带有颜色的方块，屏幕的最小组成单位；rgb 的组合；内部不会发生颜色变化；

#### 屏幕空间
屏幕坐标系，将像素的坐标用 (x, y) 表示；(2, 1) 表示左 3 下 2 的像素（下标从 0 开始）。

#### 光栅化（Rasteriza）
将多边形绘画在屏幕空间上。

### 光栅成像设备
#### CRT 显示器
阴极射线管的缩写，电子偏移并打在屏幕上成像。利用隔行扫描技术，会造成画面撕裂，产生鬼影。

#### LCD 显示器
液晶显示器。通过液晶扭曲调整光的方向。

#### LED 显示器
LED 即发光二极管，与以上显示设备原理不同。使用小的发光二极管分别成像。

### 几何基础
三角形：多边形的基础，几乎所有实体都可以拆解成三角形；可以方便地利用叉积分别内外...

### 采样
函数的离散化。一个基本的采样是判断一个像素的中心是否在三角形内。在边界上，一些软件（OpenGL 等）规定上边和左边上的点在三角形内，下边和右边上的点在三角形外。




