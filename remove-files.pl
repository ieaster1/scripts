#!/usr/bin/perl

# Super simple script to clear out files from a directory.
# main use case is in instances where rm presents "arg list too long"
# if a script is not necessary, several 1-liners can be used
# Example: perl -e 'for(<*>){unlink}'

use strict;
use warnings;
use File::Basename;

my $exec = basename($0);
my $logFile = "deleted.log";
my @files = <*>;

for my $file (@files) {
    if ( $file ne $exec ) {
        open STDOUT, '>>', $logFile;
        print "deleting $file\n";
        unlink $file or do {
            open STDERR, '>>', $logFile;
            warn("Unable to remove $file: $!");
            close(STDERR);
            next;
        };
    };
};
