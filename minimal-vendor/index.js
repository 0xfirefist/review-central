var express = require('express');
var app = express();

app.get('/', function (req, res) {
  res.json({ 
      isValid: true,
      productName: "Trimmer",
      companyName: "Philips",
      price: 5000,
      vendor: "Dummy-Vendor",
  })
});

app.listen(3000, function () {
  console.log('Vendor is listening on port 3000!');
});
