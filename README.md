## CoffeeShop Simulation

This is a basic simulator built to simulate a basic coffee shop and to play around with the Go concurrency model.

The idea of this simulator is to say how many hours a coffee shop would be opened and send customers through the coffee shop for that many hours.
While you are simulating customers going through a coffee shop, you don't want to wait for that many hours to run, the idea of a simulated day was introducted.
A simulated day allows you configure how long a day is in the simulation which allows you to express how long something should take in the simulated world like it was in real world time.

Example: Lets say you set the simulated day length to be 24 seconds long. That means that an hour is 1 second long (24 hours in day) in the simulated world.
You can then use that simulated day to express things like a minute and it'll autoamtically calculate how long a minute would be based on the simulated day length.

Once you open a coffee shop you can send customers to it at some interval for the coffee shop to accept customers and serve them.
Once the shop has closed it will no longer accept customers and will finish out serving the customers that it has in line. 

The basic idea/flow currently works like the following:

* Coffee Shop is Opened
* Baristas/Workers are spun up waiting for customers to come in
* Shop allows accepts customers and puts them in a queue
* Baristas are notified of new customers
* Baristas process the customers in the order they came in. 
* Once the bariastas are notified the shop closed, they process rest of the customers.


