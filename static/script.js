// function actuallyReady() {
//   var GoogleAuth = gapi.auth2.getAuthInstance();
//   var profile = GoogleAuth.currentUser.get();
//   if (GoogleAuth.isSignedIn.get()) {

//   console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
//   console.log('Token: ' +  profile.getAuthResponse().id_token);
//   // console.log('Name: ' + profile.getName());
//   // console.log('Image URL: ' + profile.getImageUrl());
//   // console.log('Email: ' + profile .getEmail()); // This is null if the 'email' scope is not present.	
// } else {
// 	console.log("not signed in");
// }
// }

window.onload = function() {
    gapi.load('auth2', function() {
        console.log("auth2 loaded");
        gapi.auth2.init({
            clientId: "141488749003-audttelm23ke99cmd1qgc4utd9hpqopu.apps.googleusercontent.com"
            // scope: 'some scope here'
        });
    });
};

function onSignIn(googleUser) {
  var profile = googleUser.getBasicProfile();
  document.cookie = "token=" + googleUser.getAuthResponse().id_token;
  location.reload();
}

function signOut() {
    var auth2 = gapi.auth2.getAuthInstance();
    console.log(gapi.auth2);
    auth2.signOut().then(function () {
        document.cookie = "token=";
        location.reload();
    });
}