<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        form {
            width: 200px;
            display: flex;
            flex-direction: column;
        }
        label, input, button {
            flex: 1;
            padding: 5px 0;
        }
    </style>
</head>
<body>
    <form action="http://localhost:8081/import" enctype="multipart/form-data" method="post">
        <label for="title">Image Title</label>
        <input id="title" type="text" name="title" />

        <label for="image">Upload Image File</label>
        <input id="image" type="file" name="image-file" />

        <button type="submit">Submit</button>
    </form>
</body>
</html>
