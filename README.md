# README.plz
Democratizing The Way To Help The Blind See

## Inspiration
Our friends.

We can see, but they cannot.

We can read, but they cannot.

It's **NOT** their fault that they cannot read **announcements**, **textbooks**, **ingredients list**, **menu**, and so forth. Because...they are blind.

Would it ever be possible to understand how they feel, we imagined.

What can we do to help them, we pondered.

How about we dedicate this hackathon to help our friends, we decided.

This is why we decided to work on README to increase accessibility for those who cannot see and/or read.

We hope that this project would allow our best friends from dropping out of school just because of their disabilities.

## What it does
Project README for this hackathon uses a mini-camera attached to one's glasses to retrieve text from an image and synthesizes an audio file that reads out the text for the person. Any text, handwritten or printed, can be read out loud to the person at a click of a button.

However, README's purpose is not limited to the text-reading glasses module. Its pipeline and the client application are open-source, inviting other developers to extend upon the existing project freely.
 
## How we built it

- Backend Pipeline [Publish]
  - README utilizes **Google Cloud Platform's Vision API** and **Text-To-Speech** in the backend. Using document text detection feature under GCP Vision API, we retrieve text data from an uploaded base64 image in a JSON format. Then, the text data is sent over to GCP TTS, and the service synthesizes a base64 audio file. The pipeline of the GCP APIs is written in **Golang** and is served on **Google Cloud App Engine** as a REST API.

- Frontend [Subscribe]
  - The **React.js** application, which is hosted on **Firebase**, captures a base64 image from a camera on demand and makes a POST request to the pipeline REST API. Then, once the pipeline synthesizes an audio file, the frontend application consumes the base64 data and plays back the sound on the browser.

## Challenges we ran into
Because it was our very first time using the GCP services, we weren't familiar with the client libraries. We weren't even able to download the client libraries, generate credentials file, make HTTP requests to the API, receive data, and so forth.

Also, it was our first time using Golang, and so we struggled with the basic syntax that creates a server secured on SSL/TLS. This might sound funny, but we especially struggled with understanding how *package import* works. Finally, simply uploading the Go build file to Google App Engine was a hassle as we never used GCP before.

## Accomplishments that we're proud of
Despite the challenges we ran into, we're able to figure all of them out before the submission deadline.

We're very proud that we made from zero knowledge to making an MVP prototype.

## What we learned
We learned that once we understand how one programming language works, it's easy to pick up a new language quickly.

**All we need is documentation and some example codes.**

## What's next for README
Of course, refactoring has to be done on both the backend and the frontend applications. Given that this project was created during a hackathon, we did not consider scalability and orthogonality of pragmatic programming.

On the technical side, we need to improve the accuracy of text detection regardless of the image quality. Also, we need to modulize the hardware so that it can be attached to any glasses without making it look ugly.

Finally, we acknowledge that there are many solutions like this. However, because our initial intention was ACTUALLY to help our friends and open-source the solution, we need to do the following:
- Add user-friendly control system
  - Ex. control with speech, gesture, etc
- Selective reading feature
  - Ex. if there are multiple texts, the user must be able to decide which text to read
- Non-text analysis feature
  - ex. Even if there is no text, the user must be told what's happening in front of the user.
  - ex. This new feature can inform the user of people's emotions, objects, etc
- Expand the open-source community
  - ex. recruit people to work on the open-source project to allow more people to contribute to helping disabled people

## Team:
- Alexander Bricken
- Janghoon Lee
- Alexander Wu
- Ba Thien Tran
