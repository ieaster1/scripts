#!/usr/bin/perl

# Super simple script to clear out files from a directory.
# main use case is in instances where rm presents "arg list too long"
# if a script is not necessary, several 1-liners can be used
# Example: perl -e 'for(<*>){unlink}'

use strict;
use warnings;

my $logFile = "removed.log";
my @files = <*>;

for my $file (@files) {
    open STDOUT, '>>', $logFile;
    print "removing $file\n";
    unlink $file or do {
        open STDERR, '>>', $logFile;
        warn("Unable to remove $file: $!");
        close(STDERR);
        next;
    };
};
