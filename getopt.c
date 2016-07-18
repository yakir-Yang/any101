/* getopt_demo - demonstrate getopt() usage
 *
 * This application shows you one way of using getopt() to
 * process your command-line options and store them in a
 * global structure for easy access.
 */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

/* doc2html supports the following command-line arguments:
 *
 * -I - don't produce a keyword index
 * -l lang - produce output in the specified language, lang
 * additional file names are used as input files
 *
 * The optString global tells getopt() which options we
 * support, and which options have arguments.
 */
struct globalArgs_t {
	int noIndex;				/* -I option */
	char *langCode;				/* -l option */
	const char *outFileName;		/* -o option */
	char **inputFiles;			/* input files */
	int numInputFiles;			/* # of input files */
} globalArgs;

static const char *optString = "Il:o:?";

/* Display program usage, and exit.
 */
void display_usage( void )
{
	/*
	 * noIndex       = 1
	 * langCode      = C
	 * outFileName   = output.data
	 * numInputFiles = 2
	 * inputFiles[0] = getopt.c
	 * inputFiles[1] = test.c
	 */

	puts("./out/getopt -I -l C -o output.data getopt.c test.c");
	/* ... */
	exit(EXIT_FAILURE);
}

/* Convert the input files to HTML, governed by globalArgs.
 */
void convert_document( void )
{
	printf("noIndex       = %d\n", globalArgs.noIndex);
	printf("langCode      = %s\n", globalArgs.langCode);
	printf("outFileName   = %s\n", globalArgs.outFileName);

	printf("numInputFiles = %d\n", globalArgs.numInputFiles);
	for (int i = 0; i < globalArgs.numInputFiles; i++)
		printf("inputFiles[%d] = %s\n", i, globalArgs.inputFiles[i]);

	/* ... */
}

int main( int argc, char *argv[] )
{
	int opt = 0;

	/* Initialize globalArgs before we get to work. */
	globalArgs.noIndex = 0;		/* false */
	globalArgs.langCode = NULL;
	globalArgs.outFileName = NULL;
	globalArgs.inputFiles = NULL;
	globalArgs.numInputFiles = 0;

	/* Process the arguments with getopt(), then
	 * populate globalArgs.
	 */
	while (-1 != (opt = getopt(argc, argv, optString))) {
		switch (opt) {
		case 'I':
			globalArgs.noIndex = 1;	/* true */
			break;

		case 'l':
			globalArgs.langCode = optarg;
			break;

		case 'o':
			/* This generates an "assignment from
			 * incompatible pointer type" warning that
			 * you can safely ignore.
			 */
			globalArgs.outFileName = optarg;
			break;

		case '?':
			display_usage();
			break;

		default:
			/* You won't actually get here. */
			break;
		}
	}

	globalArgs.inputFiles = argv + optind;
	globalArgs.numInputFiles = argc - optind;

	convert_document();

	return EXIT_SUCCESS;
}
