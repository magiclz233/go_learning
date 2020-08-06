FAN OUT
从同一通道读取多个功能，直到关闭该通道

Multiple functions reading from the same channel until that channel is closed

FAN IN
一个功能可以从多个输入中读取并继续进行操作，直到将所有输入通道复用到一个在关闭所有输入时都关闭的通道上，将所有通道关闭。

A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.

PATTERN

我们的管道功能有一种模式：
-当所有发送操作完成后，阶段关闭其出站通道。 -阶段将继续从入站通道接收值，直到这些通道关闭为止。

there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.