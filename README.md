Golang设计模式思想
-----
* 前言
    *   一切设计模式都是灵活应用struct的组合模式，以及go隐形继承接口的特性
    *   go中的interface就是一些方法装饰, 而struct并不依赖与接口

* 设计模式类型
  
创建模式
----
- [建造者模式(Builder Pattern)](./01-builder-patterns)
    -     将一个复杂对象的构建与它的表示分离, 使得同样的构建过程可以创建不同的表示
- [工厂方法模式(Factory Method)](./02-factory-patterns)
    -     使一个类的实例化延迟到其子类, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化
- [对象池模式(Object Pool)](./03-object-pool-pattern)
    -     根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本
- [单类模式(singleton)](./04-singleton-pattern)
    -     单例模式是最简单的设计模式之一, 保证一个类仅有一个实例, 并提供一个全局的访问接口
- [生成器(Generator)](./10-generator-pattern)
    -     生成器模式可以允许使用者在生成要使用的下一个值时与生成器并行运行
- [抽象工厂(Abstract Factory)](./11-abstract-factory)
    -     提供一个创建一系列相关或相互依赖对象的接口, 而无需指定它们具体的类
- [原型模式(Prototype Pattern)](./16-prototype-pattern)
    -     复制一个已存在的实例

结构模式
----
- [装饰模式(Decorator Pattern)](./05-decorator-pattern)
    -     装饰模式使用对象组合的方式动态改变或增加对象行为， 在原对象的基础上增加功能
- [代理模式(Proxy Pattern)](./06-proxy-pattern)
    -     代理模式用于延迟处理操作或者在进行实际操作前后对真实对象进行其它处理。
- [适配器模式(Adapter Pattern)](./12-adapter-pattern)
    -     将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
- [组合模式(Composite)](./13-composite-pattern)
    -     组合模式有助于表达数据结构, 将对象组合成树形结构以表示"部分-整体"的层次结构, 常用于树状的结构
- [享元模式(Flyweight Pattern)](./17-flyweight-pattern)
    -     把多个实例对象共同需要的数据，独立出一个享元，从而减少对象数量和节省内存
- [外观模式(Facade Pattern)]  
- [桥接模式(Bridge Pattern)]  
    
行为模式
----
- [观察者模式(Observer)](./07-observer-pattern)
    -     定义对象间的一种一对多的依赖关系,以便当一个对象的状态发生改变时,所有依赖于它的对象都得到通知并自动刷新
- [策略模式(Strategy)](./08-strategy-pattern)
    -     定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则
- [状态模式(State Pattern)](./14-state-pattern)
    -     用于系统中复杂对象的状态转换以及不同状态下行为的封装问题
- [访问者模式(Visitor Pattern)](./15-visitor-pattern)
    -     访问者模式是将对象的数据和操作分离
- [模板方法模式(Template Method Pattern)]
- [备忘录模式(Memento Pattern)]
- [中介模式(Mediator Pattern)]
- [迭代器模式(Iterator Pattern)]
- [解释器模式(Interpreter Pattern)]
- [命令模式(Command Pattern)]
- [职责链模式(Chain of Responsibility)]

同步模式(synchronization patterns)
----
- [信号量模式(Semaphore)](./09-semaphore-pattern)
    -       信号量是一种同步模式，对有限数量的资源同步互斥