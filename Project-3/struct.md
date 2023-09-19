# STRUCTS

## NOTE: We use this project to explain struct.

Strut can also be descibed as a collection of Key-value pairs (Object type in JS). 

Example:

    type car struct {
        Make string
        Model string
        Height int
        Width int
    }

we can also create a type and add it into another struct like:

    type car struct {
        Make string
        Model string
        Height int
        Width int
        FrontWheel Wheel
        BackWheel Wheel
    }

    type Wheel struct {
        Material string
        Radius int
    }

Struct can be accessed by using a 'dot' [.] operator. 

example:
Here we create a myCar variable and use the 'card' struct type and set it into a slice. 

            myCar := car{}
            myCar.FrontWheel.Radius = 5

#### Anonymous struct: The structs that doesnt have a name is called Anonymous struct. 

We only make it if there is no reason to make extra named structs type.

### Embedded structs.
Here we take one struct type and shove it into another struct.

Example:

    type car struct {
    make string
    model string
    }

    type truck struct {
    // "car" is embedded, so the definition of a
    // "truck" now also additionally contains all
    // of the fields of the car struct
    car
    bedSize int
    }

To access lets say the model, we can say truck.model instead of truck.car.model because it becomes a top model struct if we put it like this.

If we consider embedded vs nested, we should know that the embedded structs fields are accessed at the top level unlike nested structs.

NOTE: Since GO isnt an OOP, we use this method.

Another example:

    lanesTruck := truck{
    bedSize: 10,
    car: car{
        make: "toyota",
        model: "camry",
    },
    }

    fmt.Println(lanesTruck.bedSize)

    // embedded fields promoted to the top-level
    // instead of lanesTruck.car.make
    fmt.Println(lanesTruck.make)
    fmt.Println(lanesTruck.model)

## Methods in Struct

For Example:

    type rect struct {
    width int
    height int
    }

    // area has a receiver of (r rect)
    func (r rect) area() int {
    return r.width * r.height
    }

    r := rect{
    width: 5,
    height: 10,
    }

    fmt.Println(r.area())
    // prints 50

In this example we have included (r rent) as a function where we replace the struct with 'r'. 
Here we call the function 'r.area()' which returns an integer with a multiple of r.width and r.height which is stored  in r.