# Layers

### **Rabbit journey**
<small>
It will be used if the call is `async`, for exaple when we need to create an occurrence, will be sent to a queue and then the code will handle the other parts, but the client doesn't need to know what is happening
</small>

1. HTTP
2. RABBIT
3. GRPC
4. SQL


<br/><hr/><br/>


### **Default journey**
<small>
It will be used if the call is `sync`, for example when we will do the authentication, create account or other cases like those. Those cases the client must know the response, for example... the client must know that the account was created into the database and if not he need to be notified
</small>

1. HTTP
2. GRPC
3. SQL