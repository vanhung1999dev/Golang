# Epoll

epoll is a Linux-specific system call that lets you efficiently monitor many file descriptors (FDs) to see if they are ready for I/O. <br>

## ✅ How epoll works:

### 1. Create an epoll instance:

```
int epfd = epoll_create1(0);

```

### 2. Register FDs:

```
struct epoll_event ev;
ev.events = EPOLLIN;  // we want to know when it's readable
ev.data.fd = sockfd;

epoll_ctl(epfd, EPOLL_CTL_ADD, sockfd, &ev);

```

### 3. Wait for events:

```
struct epoll_event events[64];
int n = epoll_wait(epfd, events, 64, timeout_ms);

```

epoll_wait() blocks until one or more FDs are ready. It only returns ready FDs — so it's very efficient (O(1)). <br>
