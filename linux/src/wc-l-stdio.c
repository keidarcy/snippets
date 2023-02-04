#include <stdio.h>
#include <stdlib.h>

static int count(FILE *f, char *name);

int main(int argc, char *argv[])
{
	int i;
	int res = 0;
	for (i = 1; i < argc; i++)
	{
		FILE *f;
		f = fopen(argv[i], "r");

		if (!f)
		{
			exit(1);
		}

		res += count(f, argv[i]);
		fclose(f);
	}
	printf("%d total\n", res);
	exit(0);
}

int count(FILE *f, char *name)
{
	int c;
	int res = 0;
	while ((c = fgetc(f)) != EOF)
	{
		if (c == '\n')
			res++;
	}
	printf("%d %s\n", res, name);
	return res;
}