#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

static void do_tail(FILE *f, int nlines);

#define DEFAULT_N_LINES 99;

int main(int argc, char *argv[])
{
	int opt;
	int nlines = DEFAULT_N_LINES;
	int i;

	while ((getopt(argc, argv, "n:")) != -1)
	{
		switch (opt)
		{
		case 'n':
			nlines = atoi(optarg);
			break;
		case '?':
			fprintf(stderr, "bad option %s", argv[1]);
		}
	}

	if (argc == 0)
	{
		do_tail(stdin, nlines);
	}
	else
	{
		for (i = 1; i < argc; i++)
		{
			FILE *f;
			f = fopen(argv[i], "r");
			if (!f)
			{
				perror(argv[i]);
				exit(1);
			}
			do_tail(f, nlines);
			fclose(f);
		}
	}
	exit(0);
}

static void do_tail(FILE *f, int nlines)
{
	int c;

	while ((c = fgetc(f)) != EOF)
	{
		putchar(c);
	}
}
