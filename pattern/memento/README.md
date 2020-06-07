# 备忘录模式(Memento)

**定义**：

在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。

**本质**：

保持和恢复内部状态。

**优点**：

- 更好的封装性
- 简化了原发器

**缺点**：

- 如果频繁地创建备忘录开销会很大

**建议在以下情况中选用备忘录模式**：

- 如果必须保存一个对象在某一个时刻的全部或者部分状态，方便在以后需要的时候，可以把该对象恢复到先前的状态，可以使用备忘录模式。使用备忘录对象来封装和保存需要保存的内部状态，然后把备忘录对象保存到管理者对象中，在需要的时候，再从管理者对象中获取备忘录对象，来恢复对象的状态。
- 如果需要保存一一个对象的内部状态，但是如果用接口来让其他对象直接得到这些需要保存的状态，将会暴露对象的实现细节并破坏对象的封装性，这时可以使用备忘录模式，把备忘录对象实现成为原发器对象的内部类，而且还是私有的，从而保证只有原发器对象才能访问该备忘录对象。这样既保存了需要保存的状态，又不会暴露原发器对象的内部实现细节。