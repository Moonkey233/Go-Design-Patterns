# Go Design Patterns

本项目包含了使用 Go 语言实现的 22 种设计模式示例。每种设计模式都按照分类组织，并提供了相应的代码和说明。点击链接可以直接跳转到对应的代码文件。

---

## 创建型模式 (Creational Patterns)

### [抽象工厂 (Abstract Factory)](creational/abstractfactory/abstractfactory.go)
让你能创建一系列相关的对象，而无需指定其具体类。

### [生成器 (Builder)](creational/builder/builder.go)
使你能够分步骤创建复杂对象。该模式允许你使用相同的创建代码生成不同类型和形式的对象。

### [工厂方法 (Factory Method)](creational/factorymethod/factorymethod.go)
在父类中提供一个创建对象的接口，以允许子类决定实例化对象的类型。

### [原型 (Prototype)](creational/prototype/prototype.go)
让你能够复制已有对象，而又无需使代码依赖它们所属的类。

### [单例 (Singleton)](creational/singleton/singleton.go)
让你能够保证一个类只有一个实例，并提供一个访问该实例的全局节点。

---

## 结构型模式 (Structural Patterns)

### [适配器 (Adapter)](structural/adapter/adapter.go)
让接口不兼容的对象能够相互合作。

### [桥接 (Bridge)](structural/bridge/bridge.go)
可将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构，从而能在开发时分别使用。

### [组合 (Composite)](structural/composite/composite.go)
你可以使用它将对象组合成树状结构，并且能像使用独立对象一样使用它们。

### [装饰 (Decorator)](structural/decorator/decorator.go)
允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为。

### [外观 (Facade)](structural/facade/facade.go)
能为程序库、框架或其他复杂类提供一个简单的接口。

### [享元 (Flyweight)](structural/flyweight/flyweight.go)
摒弃了在每个对象中保存所有数据的方式，通过共享多个对象所共有的相同状态，让你能在有限的内存容量中载入更多对象。

### [代理 (Proxy)](structural/proxy/proxy.go)
让你能够提供对象的替代品或其占位符。代理控制着对于原对象的访问，并允许在将请求提交给对象前后进行一些处理。

---

## 行为模式 (Behavioral Patterns)

### [责任链 (Chain of Responsibility)](behavioral/chainofresponsibility/chainofresponsibility.go)
允许你将请求沿着处理者链进行发送。收到请求后，每个处理者均可对请求进行处理，或将其传递给链上的下个处理者。

### [命令 (Command)](behavioral/command/command.go)
它可将请求转换为一个包含与请求相关的所有信息的独立对象。该转换让你能根据不同的请求将方法参数化、延迟请求执行或将其放入队列中，且能实现可撤销操作。

### [迭代器 (Iterator)](behavioral/iterator/iterator.go)
让你能在不暴露集合底层表现形式（列表、栈和树等）的情况下遍历集合中所有的元素。

### [中介者 (Mediator)](behavioral/mediator/mediator.go)
能让你减少对象之间混乱无序的依赖关系。该模式会限制对象之间的直接交互，迫使它们通过一个中介者对象进行合作。

### [备忘录 (Memento)](behavioral/memento/memento.go)
允许在不暴露对象实现细节的情况下保存和恢复对象之前的状态。

### [观察者 (Observer)](behavioral/observer/observer.go)
允许你定义一种订阅机制，可在对象事件发生时通知多个“观察”该对象的其他对象。

### [状态 (State)](behavioral/state/state.go)
让你能在一个对象的内部状态变化时改变其行为，使其看上去就像改变了自身所属的类一样。

### [策略 (Strategy)](behavioral/strategy/strategy.go)
能让你定义一系列算法，并将每种算法分别放入独立的类中，以使算法的对象能够相互替换。

### [模板方法 (Template Method)](behavioral/templatemethod/templatemethod.go)
在超类中定义一个算法的框架，允许子类在不修改结构的情况下重写算法的特定步骤。

### [访问者 (Visitor)](behavioral/visitor/visitor.go)
将算法与其所作用的对象隔离开来。

---

## 使用说明

1. 克隆项目到本地：
    ```bash
    git clone git@github.com:Moonkey233/Go-Design-Patterns.git

2. 进入任意设计模式的文件夹，阅读代码示例并运行：
    ```bash
    cd ./behavioral/chainofresponsibility/
    go test
   
3. 仔细阅读代码中注释，理解设计模式的实现方式与使用场景。

---

希望本项目对您学习设计模式有所帮助！如果有任何问题，欢迎提交 Issue。

Go Design Patterns by Moonkey233 at 2024.