#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/time.h>
#include <fcntl.h>
#include <errno.h>
#include <string.h>
#include <time.h>
#include <stdarg.h>
#include <ctype.h>
#include <signal.h>

#define SERVER_NAME "LittleHTTP"
#define SERVER_VERSION "1.0"
#define HTTP_MINOR_VERSION 0
#define BLOCK_BUF_SIZE 1024
#define LINE_BUF_SIZE 4096
#define MAX_REQUEST_BODY_LENGTH (1024 * 1024)

struct HTTPHeaderField
{
	char *name;
	char *value;
	struct HTTPHeaderField *next;
};

struct HTTPRequest
{
	int protocol_minor_version;
	char *method;
	char *path;
	struct HTTPHeaderField *header;
	char *body;
	long length;
};

typedef void (*sighandler_t)(int);
static void log_exit(char *fmt, ...);
static void *xmalloc(size_t sz);
static void signal_exit(int sig);
static void install_signal_handlers(void);
static void trap_signal(int sig, sighandler_t handler);
static void install_signal_handlers(void);
static void service(FILE *in, FILE *out, char *docroot);
static void free_request(struct HTTPRequest *req);
static struct HTTPRequest *read_request(FILE *in);
static void read_request_line(struct HTTPRequest *req, FILE *in);
static struct HTTPHeaderField *read_header_field(FILE *in);
static long content_length(struct HTTPRequest *req);
static void upcase(char *str);

int main(int argc, char *argv[])
{
	if (argc != 2)
	{
		fprintf(stderr, "Usage: %s <docroot> \n", argv[0]);
		exit(1);
	}
	install_signal_handlers();
	service(stdin, stdout, argv[1]);
	exit(0);
}

static void service(FILE *in, FILE *out, char *docroot)
{
	struct HTTPRequest *req;

	req = read_request(in);
	// response_to(req, out, docroot);
	free_request(req);
}

static struct HTTPRequest *read_request(FILE *in)
{
	struct HTTPRequest *req;
	struct HTTPHeaderField *h;

	req = xmalloc(sizeof(struct HTTPRequest));
	read_request_line(req, in);
	req->header = NULL;
	while ((h = read_header_field(in)))
	{
		h->next = req->header;
		req->header = h;
	}

	req->length = content_length(req);
	if (req->length != 0)
	{
		if (req->length > MAX_REQUEST_BODY_LENGTH)
		{
			log_exit("request body too long");
		}
		req->body = xmalloc(req->length);
		if (fread(req->body, req->length, 1, in) < 1)
		{
			log_exit("failed to read request body");
		}
	}
	else
	{
		req->body = NULL;
	}
	return req;
}

static void read_request_line(struct HTTPRequest *req, FILE *in)
{
	char buf[LINE_BUF_SIZE];
	char *path, *p;

	if (!fgets(buf, LINE_BUF_SIZE, in))
	{
		log_exit("no request line");
	}

	p = strchr(buf, ' ');
	if (!p)
	{
		log_exit("parse error on request line (1): %s", buf);
	}
	*p++ = '\0';

	req->method = xmalloc(p - buf);
	strcpy(req->method, buf);
	upcase(req->method);

	path = p;
	p = strchr(path, ' ');
	if (!p)
	{
		log_exit("parse error on request line (2): %s", buf);
	}
	*p++ = '\0';
	req->path = xmalloc(p - path);
	strcpy(req->path, path);

	if (strncasecmp(p, "HTTP/1.", strlen("HTTP/1.")) != 0)
	{
		log_exit("parse error on request line(3): %s", buf);
	}
	p += strlen("HTTP/1.");
	req->protocol_minor_version = atoi(p);
}

static struct HTTPHeaderField *read_header_field(FILE *in)
{
	struct HTTPHeaderField *h;
	return h;
}

static long content_length(struct HTTPRequest *req)
{
	long len;
	return len;
}

static void
free_request(struct HTTPRequest *req)
{
	struct HTTPHeaderField *h, *head;

	head = req->header;
	while (head)
	{
		h = head;
		head = head->next;
		free(h->name);
		free(h->value);
		free(h);
	}
	free(req->method);
	free(req->path);
	free(req->body);
	free(req);
}

static void log_exit(char *fmt, ...)
{
	va_list ap;
	va_start(ap, fmt);
	vfprintf(stderr, fmt, ap);
	fputc('\n', stderr);
	va_end(ap);
	exit(1);
}

static void *xmalloc(size_t sz)
{
	void *p;

	p = malloc(sz);
	if (!p)
	{
		log_exit("failed to allocate memory");
	}
	return p;
}

static void install_signal_handlers(void)
{
	trap_signal(SIGPIPE, signal_exit);
}

static void trap_signal(int sig, sighandler_t handler)
{
	struct sigaction act;

	act.sa_handler = handler;
	sigemptyset(&act.sa_mask);
	act.sa_flags = SA_RESTART;
	if (sigaction(sig, &act, NULL) < 0)
	{
		log_exit("sigaction() failed: %s", strerror(errno));
	}
}

static void signal_exit(int sig)
{
	log_exit("exit by signal %d", sig);
}

static void upcase(char *str)
{
	char *p;

	for (p = str; *p; p++)
	{
		*p = (char)toupper((int)*p);
	}
}