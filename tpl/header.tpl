{{define "header"}}
<!DOCTYPE html>
<html lang="en">
<head>

  <!-- Basic Page Needs
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <meta charset="utf-8">
  <title>SerekeMail</title>
  <meta name="description" content="Sereke is a Modern Web Mail API">
  <meta name="author" content="Ivan Chavero">

  <!-- Mobile Specific Metas
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <!-- FONT
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <link href="//fonts.googleapis.com/css?family=Raleway:400,300,600" rel="stylesheet" type="text/css">

  <!-- CSS
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <link rel="stylesheet" href="/static/css/sereke.css">
  <link rel="stylesheet" href="/static/skeleton/css/normalize.css">
  <link rel="stylesheet" href="/static/skeleton/css/skeleton.css">

  <!-- JAVASCRIPT
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <script src="/static/js/jquery-3.1.1.js"></script>

  <!-- Favicon
  –––––––––––––––––––––––––––––––––––––––––––––––––– -->
  <link rel="icon" type="image/png" href="/skeleton/images/favicon.png">

  <script language="javascript">
    function getLayout()

    $.get("/folders").then(function(data) {
        $(".folders").html(data);
    });
  </script>

</head>
<body onLoad="getLayout">
    <div class="header">
        <h1>S U P E R &nbsp;&nbsp; S E R E K E &nbsp;&nbsp; W E B M A I L &nbsp;&nbsp; A P I </h1>
        <div class="mail_controls">
            <button>&lt;&lt;Prev</button> &nbsp;&nbsp; <button>Next&gt;&gt;</button>
        </div>
    </div>
{{end}}
