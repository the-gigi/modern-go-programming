# Generic Programming in Go

## Built-in generic containers

### The empty interface
- An interface is a pointer
- Objects that operate at the generic level
- Using an empty interface
- Interfaces with empty interface arguments and specialized implementations
	
## Code Generation of specialized containers

### Build your own set
- Set of empty interfaces
- Set with generic interface and specialized implementations
    - The Set interface accepts empty interface
    - No type-safety
    - Need to cast

### Generating strongly-typed sets
- All the benefits
- Need to use a code generator
