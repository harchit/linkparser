Parses an HTML doc tree and obtains a list of <a> tags

<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/harchit_b">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/harchit">
      I am on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>

returns:

[{Href:https://www.twitter.com/harchit_b Text:Check me out on twitter} {Href:https://github.com/harchit Text:I am on Github !}]
