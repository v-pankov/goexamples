Usecase package

For now the only two exercises user can do with this application is to **login** and **create** messages.

The core of the app does not bother what happens with created messages afterwards. Created messages can
be polled by the client, or they can be broadcasted to every client by the server. But this is not a concern
for the core. Physical delivery of created message does not have any business value so it does not deserve a
usecase. Bringing such usecases up increases complexity of the system. I may be wrong but that's where I came
up so far.

This is a note for the future me. I must read it before re-doing eveything.

---

I'll leave this question here: can presenter be asynchronous?

For example, in this particular project user creates a message, the message is than being put to database and
pushed to event bus. After some time this event is received somewhere and it must be sent to the client.

So can presenter and usecase live in different threads of execution? It seems good if true but there is an obstacle:
presenter shares usecase response model but event received from the bus does not have to have such structure.

---

I think the best solution would be to have actual usecase for processing events from the bus. It could save
interesting events to event log and do some other analytics work and than let presenter to send message to
client.

Right now there is no need for such usecase. So it's better just to send message to the client once event is
received without using special usecase for that.