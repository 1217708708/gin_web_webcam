$(() => {
    Webcam.set({
        width: 320,
        height: 240,
        dest_width: 640,
        dest_height: 480,
        image_format: 'jpeg',
    });
    Webcam.attach('#my_camera');
    console.log("ok");
    Webcam.on( 'load', function() {
        // library is loaded
        test()
    } );
})
function test() {
    Webcam.snap(function (data_uri) {
        // do something with the data URI!
        const image = new Image();
        image.src = data_uri;
        // document.getElementById('my_camera').appendChild(image);
        // post data_uri to server for save image
        // console.log(image.src);
        $.ajax({
            url: "/service/upload",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify({
                "data": data_uri
            }),
            processData: false,
            success: function (data) {
                console.log("Upload success.");
            },
            error: function (err) {
                console.log("Upload failed.");
            }
        });
    });
}
