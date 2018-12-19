# A Simple BDD Demo with GoConvey, and Beego
## Introduction

Im creating a simple ```e-commerce``` api where user can login and purchase a specific product.

Me as a product owner creating a several user stories, both positive and negative flow .

### Story 1 
```
GIVEN Edwin as a USER 
    WHEN Edwin login with the correct username and password 
      THEN HTTP Code Should Be 200 .
      AND The Result Should Not Be Empty .
      AND The Result Should Contains "login success" .
```

### Story 2 
```
GIVEN Edwin as a USER 
    WHEN Edwin login with the WRONG username and password 
      THEN HTTP Code Should Be 404 .
      AND The Result Should Not Be Empty .
      AND The Result Should Contains "user not exist" .
```