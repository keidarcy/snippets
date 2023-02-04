#include <stdio.h>
#include <stdlib.h>

static void do_cat(const char *path);
static void die(const char *path);

int main(int argc, char *argv[])
{
	if (argc < 0)
	{
		fprintf(stderr, "%s: file name not given\n", argv[0]);
	}
	int i;
	for (i = 1; i < argc; i++)
	{
		do_cat(argv[i]);
	}
	exit(0);
}

static void do_cat(const char *path)
{
	unsigned char buf[BUFSIZ];
	FILE *f = fopen(path, "r");
	if (!f)
		die(path);

	for (;;)
	{
		size_t n_read = fread(buf, 1, sizeof buf, f);
		if (ferror(f))
			die(path);
		size_t n_written = fwrite(buf, 1, n_read, stdout);
		if (n_written < n_read)
			die(path);
		if (n_read < sizeof buf)
			break;
	}
	fclose(f);
}

static void die(const char *s)
{
	perror(s);
	exit(1);
}