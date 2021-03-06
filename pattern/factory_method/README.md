# 工厂方法模式(Factory Method)

工厂方法模式很好地体现了“**依赖倒置原则**”。

依赖倒置原则告诉我们“**要依赖抽象，不要依赖于具体实现**”。简单点来说就是：不能让高层组件依赖于低层组件，而不管高层组件还是底层组件，都应该依赖于抽象。

比如对于一个客户端想要的导出的操作`ExportOperate`就是高层组件，而具体实现数据导出的对象就是低层组件，比如导出文本文件`ExportTextFile`、导出到数据库`ExportDB`，而抽象就是导出操作的接口`ExportApi`。

**定义**：

定义一个用于创建对象的接口，让子类决定实例化哪一个类，工厂方法使一个类的实例化延迟到其子类。

**本质**：

延迟到子类来选择实现。

**优点**：

- 可以在不知具体实现的情况下编程
- 更容易扩展对象的新版本

**建议在以下情况中选用工厂方法模式**：

- 如果一个类需要创建某个接口的对象，但又不知道具体的实现，这种情况可以选用工厂方法模式，把创建对象的工作延迟到子类中去实现。
- 如果一个类本身就希望由它的子类来创建所需的对象的时候，应该使用工厂方法模式。