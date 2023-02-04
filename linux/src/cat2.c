#include <stdio.h>
#include <stdlib.h>

static void do_cat(FILE *f);

int main(int argc, char *argv[])
{
	int i;

	for (i = 1; i < argc; i++)
	{
		FILE *f;
		f = fopen(argv[i], "r");
		if (!f)
			exit(1);
		do_cat(f);
		fclose(f);
	}
	exit(0);
}

static void do_cat(FILE *f)
{
	int c;
	while ((c = fgetc(f)) != EOF)
	{
		if (putchar(c) < 0)
			exit(1);
	}
}