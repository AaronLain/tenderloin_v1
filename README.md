# Tenderloin

## A specific solution to a very specific problem

Tenderloin is a simple api to aid in the shipment of perishable product. It takes in a csv file of zip codes and assigns an ice profile for shipping based on the highest temperature of the destination zip during the transit window.

To install: make sure you have golang installed, then clone this repo. You will have to install gocsv (a future update might forgo this, but it works for now) 

```
go get -u github.com/gocarina/gocsv
```

Then simply build the `main.go` package 
```
go build main.go
```

Now all you have to do is run the executable with the path of your source csv file.

```
./main /path/to/your/csv
```

Step 4: Profit.
