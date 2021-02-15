- [Store Pickup by Secomapp](#store-pickup-by-secomapp)
- [Bespoke_Shipping](#bespoke_shipping)
	- [instruction](#instruction)
- [test script](#test-script)
-
## Store Pickup by Secomapp

- [Store Pickup by Secomapp](https://apps.shopify.com/store-pickup-3)

- [target](https://prnt.sc/vvmz4e)

## Bespoke_Shipping

-[Bespoke Shipping](https://apps.shopify.com/custom-shipping-rates)

### instruction

> existing sample

```php
/* This macro will be parsed as PHP code (see http://www.php.net)

The calculateshipping function is called every time a shipping calculation request is made by Shopify.

The function must return an array of available shipping options, otherwise no shipping options will be returned to your customers.
*/
function calculateshipping($DATA) {

	/* do not edit above this line */

	/*
	NB: only select php functions are allowed. Use of anything else will return a syntax error when you try to save. Consult the user guide at http://www.parcelintelligence.com.au/cs/documentation for more information
	Examples to get you started are available here: http://parcelintelligence.com.au/cs/documentation/examples/
	To take advantage of our assisted setup option, please email hello@parcelintelligence.com.au with a detailed description of how you want your shipping rates setup for a quote.
	*/

	//this holds the rates that will be returned to your customer
	$_RATES = array();

	/*
	this is what $DATA looks like, you can use any of the info in $DATA to generate your shipping rate
	use print_r($DATA); to see actual contents of the $DATA array in the logs.
	Array
	(
		[origin] => Array
			(
				[country] => AU
				[postal_code] => 3000
				[province] => VIC
				[city] => melbourne
				[name] =>
				[address1] => 1 main street
				[address2] =>
				[address3] =>
				[phone] =>
				[fax] =>
				[address_type] =>
				[company_name] =>
			)

		[destination] => Array
			(
				[country] => AU
				[postal_code] => 2000
				[province] => NSW
				[city] =>
				[name] =>
				[address1] =>
				[address2] =>
				[address3] =>
				[phone] =>
				[fax] =>
				[address_type] =>
				[company_name] =>
			)

		[items] => Array
			(
				[0] => Array
					(
						[name] => product10
						[sku] => SKUP00310
						[quantity] => 1
						[grams] => 1000
						[price] => 300
						[vendor] => PRODUCT
						[requires_shipping] => 1
						[taxable] => 1
						[fulfillment_service] => manual
						[product_id] => 128436738
						[variant_id] => 290813760
					)

				[1] => Array
					(
						[name] => product11
						[sku] => SKUP0011
						[quantity] => 1
						[grams] => 1100
						[price] => 300
						[vendor] => PRODUCT
						[requires_shipping] => 1
						[taxable] => 1
						[fulfillment_service] => manual
						[product_id] => 128436744
						[variant_id] => 290813772
					)

			)

		[currency] => AUD
	)

	//this is how you insert a rate
	$_RATES[] = array(
		"service_name" => "Standard Shipping", //this is what the customer will see
		"service_code" => "STANDARD_SHIPPING", //can be anything you like
		"total_price" => 10000, //in cents
		"currency" => "AUD",
	);
	*/

	return $_RATES;

	/* do not edit below this line */

}
```

## test script

original price * 100 = *yen\* price

```php
$DATA = ['items' => [['quantity' => 12,'grams' => 1300,'price' => 600000]]];

$totalPrice = 0;
$numberPerBox = 12;
$boxPrice = 1000 + 800;
$pricePerTime = 800;
$pricePerWeight = 910;
$boxWeight = 3.5;
$isDDP = true;
$rateDDP = 1.25;
$priceDDP = 4000;
$priceFrance = 5000;


$totalWeight = 0;
$boxNumber = 0;
$productPrice = 0;
foreach ($DATA['items'] as $index => $item) {
    $totalWeight += $item['quantity'] * ($item['grams'] / 1000);
    $boxNumber += ceil($item['quantity'] / $numberPerBox);
    $productPrice += $item['quantity'] * ($item['price'] / 100);
}
$totalWeight += $boxNumber * $boxWeight;
$boxPrice = ($boxNumber * $boxPrice) + $pricePerTime;
$shippingPrice = $totalWeight * $pricePerWeight;


if ($isDDP) {
    $totalPrice = floor((($shippingPrice + $boxPrice + $productPrice) * $rateDDP) + $priceDDP);
} else {
    $totalPrice = floor(($shippingPrice + $boxPrice + $productPrice));
}

if ($DATA['destination']['country'] == 'FR') {
    $totalShippingPrice += $priceFrance;
}

$totalShippingPrice = ($totalPrice - $productPrice) * 100;

$_RATES[] = array(
	"service_name" => "Standard Shipping", //this is what the customer will see
	"service_code" => "STANDARD_SHIPPING", //can be anything you like
	"total_price" => $totalShippingPrice, //in cents
	"currency" => "JPY",
);

var_dump($productPrice);
var_dump($shippingPrice);
var_dump($boxPrice);
var_dump($totalPrice);
var_dump($totalShippingPrice);
```
