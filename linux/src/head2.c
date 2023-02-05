#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

static void do_head(FILE *f, int nlines);
#define DEFAULT_N_LINES 10;
int main(int argc, char *argv[])
{
	int opt;
	long nlines = DEFAULT_N_LINES;

	while ((opt = getopt(argc, argv, "n:")) != -1)
	{
		switch (opt)
		{
		case 'n':
			nlines = atol(optarg);
			break;
		case '?':
			fprintf(stderr, "Usage: %s n [file file...]\n", argv[0]);
			exit(1);
		}
	}

	if (optind == argc)
	{
		do_head(stdin, nlines);
	}
	else
	{
		int i;
		for (i = optind; i < argc; i++)
		{

			FILE *f;
			f = fopen(argv[i], "r");
			if (!f)
			{
				perror(argv[i]);
				exit(1);
			}
			do_head(f, nlines);
			fclose(f);
		}
	}
	exit(0);
}

static void do_head(FILE *f, int nlines)
{
	int c;
	while ((c = fgetc(f)) != EOF)
	{
		if (putchar(c) < 0)
		{
			exit(1);
		}

		if (c == '\n')
		{
			nlines--;
			if (nlines == 0)
			{
				return;
			}
		}
	}
}