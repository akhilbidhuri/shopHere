<h2>ShopHere</h2>

<p>Online store providing following features-</p>

<ul>
<li>/user/create - Creates a new User</li>
<li>/user/login - Login for existing user based on username and
password</li>
<li>/item/create - Creates an Item</li>
<li>/cart/add* - Adds Items to the cart</li>
<li>/cart/:cartId/complete* - Converted the cart into an order</li>
<li>/user/list - List all users</li>
<li>/item/list - List all items</li>
<li>/cart/list - List all carts</li>
<li>/order/list - List all orders</li>
<p>* uses token</p>
</ul>

Provides single device support only one active token can exist for a user.

Tech Stack - Go, Gin, Gorm, Postgresql