# 状态模式(State)

**定义**：

允许一个对象在其内部状态改变时改变它的行为。对象看起来似乎修改了它的类。

**本质**：

根据状态来分离和选择行为。

**优点**：

- 简化应用逻辑控制
- 更好地分离状态和行为
- 更好的扩展性
- 显式进行状态转换

**缺点**：

- 一个状态对应一个状态处理类，会使得程序引入太多的状态类，这让程序变得混乱

**建议在以下情况中选用状态模式**：

- 如果一个对象的行为取决于它的状态，而且它必须在运行时刻根据状态来改变它的行为，可以使用状态模式，来把状态和行为分离开。虽然分离开了，但状态和行为是有对应关系的，可以在运行期间，通过改变状态，就能够调用到该状态对应的状态处理对象上去，从而改变对象的行为。
- 如果一个操作中含有庞大的多分支语句，而且这些分支依赖于该对象的状态，可以使用状态模式，把各个分支的处理分散包装到单独的对象处理类中，这样，这些分支对应的对象就可以不依赖于其他对象而独立变化了。

[跳转到implement](../../implement/state)