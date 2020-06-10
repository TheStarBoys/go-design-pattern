# go-design-pattern
go设计模式，研磨设计模式的读书笔记。

## 简介

### 概念

设计模式：是指在软件开发中，经过验证的，用于解决在**特定环境**下**重复出现**的**特定问题**的**解决方案**。

### 学习理由

1. 设计模式已经成为软件开发人员的“标准词汇”
2. 学习设计模式是个人技术能力提高的捷径
3. 不用重复发明轮子

### 学习设计模式的层次

1. 基本入门级
   - 要求能够正确理解和掌握每个设计模式的基本知识，能够识别在什么场景下出现了什么样的问题，采用何种方案来解决它，并能够在实际的程序设计和开发中套用相应的设计模式。
2. 基本掌握级
   - 除了具备基本入门级的要求外，还要求能够结合实际应用的场景，对设计模式进行变形使用。
   - 能在实际开发中，根据需求，对模式进行变形。变形的前提是要能准确深入地理解和把握设计模式的本质，**万变不离其宗，只有把握本质，才能确保正确变形使用而不是误用**。
3. 深入理解和掌握级
   - 从思想和方法上吸收设计模式精髓，**收发随心**。
   - 根据需求，对已有的模式进行**变形**、**组合**。

## 项目结构

`./pattern` 包里放置比较通用的模板

`./implement` 包里放置相关模式的具体实现，比较简单的设计模式没有单独在此包进行实现。



## 分类

### 创建型模式

抽象了对象实例化过程，用来帮助创建对象的实例。

- [简单工厂模式 (Simple Factory)](./pattern/simple_factory)
- [工厂方法模式 (Factory Method)](./pattern/factory_method)
- [抽象工厂模式 (Abstract Factory)](./pattern/abstract_factory)
- [生成器模式 (Builder)](./pattern/builder)
- [原型模式 (Prototype)](./pattern/prototype)
- [单例模式 (Singleton)](./pattern/singleton)

### 结构型模式

描述如何组合类和对象以获得更大的结构。

- [外观模式 (Facade)](./pattern/facade)
- [适配器模式 (Adapter)](./pattern/adapter)
- [代理模式 (Proxy)](./pattern/proxy)
- [组合模式 (Composite)](./pattern/composite)
- [享元模式 (Flyweight)](./pattern/flyweight)
- [装饰模式 (Decorator)]()
- [桥接模式 (Bridge)]()

### 行为型模式

描述算法和对象间职责的分配。

- [中介者模式 (Mediator)](./pattern/mediator)
- [观察者模式 (Observer)](./pattern/observer)
- [命令模式 (Command)](./pattern/command)
- [迭代器模式 (Iterator)](./pattern/iterator)
- [模板方法模式 (Template Method)](./pattern/template_method)
- [策略模式 (Strategy)](./pattern/strategy)
- [状态模式 (State)](./pattern/state)
- [备忘录模式 (Memento)](./pattern/memento)
- [解释器模式 (Interpreter)]()
- [职责链模式 (Chain of Responsibility)]()
- [访问者模式 (Visitor)]()