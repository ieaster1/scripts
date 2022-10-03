#!/usr/bin/perl

use strict;
use warnings;

my @files = <*>;

foreach my $file (@files) {
    open STDOUT, '>>', "remove.log";
    print "removing $file\n";
    unlink $file or die "Unable to unlink $file: $!";
}
