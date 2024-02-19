# In Memory Key Value Store and a Queue Store
This application imitates an in memory key value store with help of golang.
It accepts multiple commands such as SET, GET, QPUSH ,QPOP , BQPOP to perform various operations on the store

## Running locally
This application can be run locally in 2 ways
1. Using Golang Compiler
2. Using Docker

### Requirments
1. Golang [How to install Golang ? ](https://go.dev/doc/install)
2. Docker [How to install Docker ? ](https://docs.docker.com/engine/install/)

#### 1. Using Golang Compiler
1. Clone the repository and open the root folder.
2. Run ``` go build -o main ```
3. To start the application run ``` ./main ```
4. The server is up and running on port 8080.

#### 2. Using Docker
1. Clone the repository and open the root folder.
2. Run ``` docker build -t inmemorystore . ``` to build the docker image.
3. To start the application run ``` docker run -it -p 8080:8080 --rm --name inmemstore inmemorystore ``` .
4. The server is up and running on port 8080.

## Usage
### 1. Key Value Store
#### 1.1 Pattern: 
`SET <key> <value> <expiry time>? <condition>?`

- **`<key>`**:  
  The key under which the given value will be stored.

- **`<value>`**:  
  The value to be stored.

- **`<expiry time>`**:  
  Specifies the expiry time of the key in seconds.  
  Must contain the prefix `EX`.  
  This is an optional field.  
  The field must be an integer value.

- **`<condition>`**:  
  Specifies the decision to take if the key already exists.  
  Accepts either `NX` or `XX`.  
  - `NX` -- Only set the key if it does not already exist.  
  - `XX` -- Only set the key if it already exists.  
  This is an optional field.  
  The default behavior will be to upsert the value of the key.

#### Examples of SET Command:

- **`SET key_a 2`**:  
  Sets the value 2 in `key_a` and does not expire.

- **`SET key_b 3 EX 60`**:  
  Sets the value 3 in `key_b` and expires in 60 seconds.

- **`SET key_c 4 EX 60 NX`**:  
  Sets the value 4 in `key_c`, expires in 60 seconds, and only sets the value if the key does not already exist.

- **`SET key_d 5 XX`**:  
  Sets the value 5 in `key_d`, does not expire, and only sets the value if the key already exists.

#### 1.2 Pattern: 
`GET <key>`

- **Examples**:
  - `GET key_a`:  
    Returns the value stored using the specified key.


### 2. Queue
#### 2.1 Pattern: 
`QPUSH <key> <value...>`

- **`<key>`**:  
  Name of the queue to write to.

- **`<value...>`**:  
  Variadic input that receives multiple values separated by space.

- **Examples**:
  - `QPUSH list_a 1`:  
    Adds value `1` to the queue named `list_a`.
  - `QPUSH list_a 2 3 4`:  
    Adds values `2`, `3`, and `4` to the queue named `list_a`.

#### 2.2 Pattern: 
`QPOP <key>`

- **`<Key>`**:  
  Name of the queue.

- **Examples**:
  - `QPUSH list_a 1`:  
    Returns OK.
  - `QPOP list_a`:  
    Returns `1`.
  - `QPUSH list_a 1 2`:  
    Returns OK.
  - `QPOP list_a`:  
    Returns `1`.
  - `QPOP list_a`:  
    Returns `2`.
  - `QPOP list_x`:  
    Returns `null`.

### 3. Blocking Queue
#### 3.1 Pattern: 
`BQPOP <key> <timeout>`

- **`<key>`**:  
  Name of the queue to read from.

- **`<timeout>`**:  
  The duration in seconds to wait until a value is read from the queue.  
  The argument must be interpreted as a double value.  
  A value of `0` immediately returns a value from the queue without blocking.

- **Example**:
  - **Scenario 1**:  
    - `QPUSH list_1 a`:  
      Returns OK.
    - `BQPOP list_1 0`:  
      Returns `a`.

  - **Scenario 2**:  
    - `BQPOP list_1 0`:  
      Returns `null`.

  - **Scenario 3**:  
    - `BQPOP list_1 10`:  
      Blocks the request as the queue is empty.  
      Returns `null` after 10 seconds.


