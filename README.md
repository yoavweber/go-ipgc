# PGFS - ParaGliding File System
## Project Description
This project is built as (self)tasked in the course PROG2005 Cloud Technologies for the Spring 2021 Semester. The project is suggested to be available as open source, for further work from external contributors after the project period has ended.
This application is based IPFS (InterPlanetary File System), and is intended to be used by paragliders to share flight-info with each other, completely decentralized. The application in the current state is entirely backend-focused and is written in golang.
Paraglider-flight files (in standardized format ".igc") are intended for sharing on a peer-to-peer basis.

This project is directly forked from the go-implementation of IPFS - located at [go-ipfs](https://github.com/ipfs/go-ipfs)<br>
All licenses and dependencies from the original project has been preserved.

Further Project Progress Documentation is located further down in this README.

## Features
The program is directly based on IPFS, and all features from the go-implementation of IPFS ([go-ipfs](https://github.com/ipfs/go-ipfs)) is still operational. Additional features (as described in the project description) include:
* Uploading and downloading files is exclusive for files of the .igc-format (standardized file-format for paragliding flight-information)
* Running the program initiates 3 PGFS-nodes for testing-purposes
* Mother-node (main server bootstrap-node) is located on the internal NTNU-network. In practice, this means the program only works internally in the NTNU institutional network. To use your own server-node (for further development purposes), the bootstrap-node must be updated to coincide with your own server-node.

## Deployment
This project is planned to be dockerized closer to due-date. More information coming soon!

## The Development Process (for assessment purposes)
### Original Project Plan
* Brief description - may be redundant, as it is mostly stated in the project description and feature-set.
### What went well
* Here goes what has been implemented, what has the group learned from the experience?
### What didn't go as expected
* Here goes issues - what was hard (getting to know IPFS on a deep level was both challenging and rewarding + took a lot of time (way more than expected in the proposed timeframe), write a paragraph or two on this)
* Some platform issues, we are a diverse group in terms of preferred OS. 1x Linux, 1x Windows, 1x macOS.
### The most challenging aspects
* Here goes aspects of the project the group deemed most challenging.
### The learning outcome
* What did we learn? IPFS being the obvious answer, but it may be worth adding stuff like interrogating a foreign repository to make it fit specialized needs.
### Total work hours
For this project, a high focus on collaboration has been prioritized. For every week of the project period, the group has had 6-10 hours of collaborative work (usually 2-3 sessions per week). The sessions have mostly happened physically (on campus) and each session usually has had a specific purpose (some sprints, some commune problem solving). In addition to these teamwork-sessions, we also split some work, to increase the learning-outcome per individual group member. ----put some more later, but ground work has been laid ;)



<br>Authored by<br>
<b>Yoav Weber</b><br>
<b>Milosz Antoni Wudarczyk</b><br>
<b>Kristian Amundsen Øhman-Norén</b><br>
2021, Norwegian University of Science and Technology
