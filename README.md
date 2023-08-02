# TK4CTL

## Synopsis
tk4ctl is a helper utility to manage a Hercules MVS instance as well as the guest OS loaded into the emulator.
 
## Assumptions
* There is a need for the solution.
* It's possible to run ultimately any Hercules command.
* It's possible to run ultimately any z/OS command.
* The Hercules CGI API will be used.
* Configurations may or may not change.
* The utility may or may not need to be ran remotely.
* Either HTTP or TCP will be used.
* The utility will be written in Go.
* The utility will be ran on POSIX systems.
 
## Requirements
* The utility will be able to run commands through the Hercules HTTP endpoints.
* The utility will be able to retrieve the system log.
* The utility will be able to detect both guest and host OS's.
* The utility will be able to map common command names to z/OS commands.
* The utility will be able to generate JCL job cards.
* The utility will be able to submit JCL job cards.
 
## Design
### Operations
Operation    | Description                                         |
|--|--|
|tso         | Runs a raw TSO command                              |
|jcl	        | Runs a JCL job via sockdev                          |
|herc        | Runs a raw Hercules command through HTTP            |
|kicks       | Executes a KICKS process (assumes KICKS installed)  |

### Configuration
|Keys              | Description                                         |
|--|--|
|TraceLevel	       | Level of the tracing.  Possible values are 0-4.     |

### Flags
All commands have the format tk4ctl <operation> <flags>
|Flag            | Description                                      |
|--|--|                                                                                                                                                                                                     |-h	         | ip:port combination                                  | 
|-c	         | Base64 encoded MVS credentials to run JCL jobs under |
|-m	         | MVS dataset member                                   |
|-f	         | Local file to run                                    |
|-e	         | Command arguments                                    |
|--version	  | Prints the version of the utility                    |
|--help	     | Prints the help menu                                 |
Note that:
* Local file is local to the client
* Dataset and file are mutually exclusive, if both are provided, local file will be used
* For tso commands, it is required to redirect the job output to httproot by changing the following lines:
```bash
   # 000E 1403 prt/prt00e.txt ${TK4CRLF}
   000E 1403 hercules/httproot/tk4ctl_log.txt ${TK4CRLF}
```

## Examples
```bash
./dist/tk4ctl jcl -h "localhost:8038" -c "SEVSQzAxOkNVTDhUUg==" -f ./MINMAX_COBOL 
\n14.22.24           $HASP160 PRINTER1 INACTIVE - CLASS=A\n14.22.24 STC  334  IRB101I MF/1 REPORT AVAILABLE FOR PRINTING\nHHC01040I 0:000C COMM: client &lt;unknown&gt;, ip 172.17.0.1 connected to device 3505\nHHC01206I 0:000C Card: client &lt;unknown&gt;, ip 172.17.0.1 disconnected from device 3505\n14.32.21 JOB   61  $HASP100 HERC02A  ON READER1     MIN AND MAX\n14.32.21 JOB   61  $HASP373 HERC02A  STARTED - INIT  1 - CLASS A - SYS TK4-\n14.32.21 JOB   61  IEF403I HERC02A - STARTED - TIME=14.32.21\n14.32.22 JOB   61  IEF404I HERC02A - ENDED - TIME=14.32.22\n14.32.22 JOB   61  $HASP395 HERC02A  ENDED\n14.32.22           $HASP309    INIT  1 INACTIVE ******** C=A\n14.37.24 STC  334  $HASP150 MF1      ON PRINTER1       208 LINES\n14.37.24           $HASP160 PRINTER1 INACTIVE - CLASS=A\n14.37.24 STC  334  IRB101I MF/1 REPORT AVAILABLE FOR PRINTING\n14.37.25           $HASP000 OK\nHHC01040I 0:000C COMM: client &lt;unknown&gt;, ip 172.17.0.1 connected to device 3505\nHHC01206I 0:000C Card: client &lt;unknown&gt;, ip 172.17.0.1 disconnected from device 3505\n14.40.00 JOB   62  $HASP100 HERC02A  ON READER1     MIN AND MAX\n14.40.00 JOB   62  $HASP373 HERC02A  STARTED - INIT  1 - CLASS A - SYS TK4-\n14.40.00 JOB   62  IEF403I HERC02A - STARTED - TIME=14.40.00\n14.40.01 JOB   62  IEF404I HERC02A - ENDED - TIME=14.40.01\n14.40.01 JOB   62  $HASP395 HERC02A  ENDED\n14.40.01           $HASP309    INIT  1 INACTIVE ******** C=A\nHHC01040I 0:000C COMM: client &lt;unknown&gt;, ip 172.17.0.1 connected to device 3505\nHHC01206I 0:000C Card: client &lt;unknown&gt;, ip 172.17.0.1 disconnected from device 3505\n14.45.13 JOB   63  $HASP100 HERC02A  ON READER1     MIN AND MAX\n14.45.13 JOB   63  $HASP373 HERC02A  STARTED - INIT  1 - CLASS A - SYS TK4-\n14.45.13 JOB   63  IEF403I HERC02A - STARTED - TIME=14.45.13\n14.45.14 JOB   63  IEF404I HERC02A - ENDED - TIME=14.45.14\n14.45.14 JOB   63  $HASP395 HERC02A  ENDED\n14.45.14           $HASP309    INIT  1 INACTIVE ******** C=A\n

./dist/tk4ctl tso -h "localhost:8038" -c "SEVSQzAxOkNVTDhUUg==" -e "cmd=date"

IEF376I  JOB /HERC01G / STOP  23210.1445 CPU    0MIN 00.03SEC SRB    0MIN 00.00SEC

 date
DATE = 07/29/23  (23.210)    TIME = 14.45.25
END
                        HH        HH  EEEEEEEEEEEE  RRRRRRRRRRR    CCCCCCCCCC     00000000         11        GGGGGGGGGG
                       HH        HH  EEEEEEEEEEEE  RRRRRRRRRRRR  CCCCCCCCCCCC   0000000000       111       GGGGGGGGGGGG
                      HH        HH  EE            RR        RR  CC        CC  00      0000     1111       GG        GG
                     HH        HH  EE            RR        RR  CC            00     00 00       11       GG
                    HH        HH  EE            RR        RR  CC            00    00  00       11       GG
                   HHHHHHHHHHHH  EEEEEEEE      RRRRRRRRRRRR  CC            00   00   00       11       GG
                  HHHHHHHHHHHH  EEEEEEEE      RRRRRRRRRRR   CC            00  00    00       11       GG     GGGGG
                 HH        HH  EE            RR    RR      CC            00 00     00       11       GG     GGGGG
                HH        HH  EE            RR     RR     CC            0000      00       11       GG        GG
               HH        HH  EE            RR      RR    CC        CC  000       00       11       GG        GG
              HH        HH  EEEEEEEEEEEE  RR       RR   CCCCCCCCCCCC   0000000000    1111111111   GGGGGGGGGGGG
             HH        HH  EEEEEEEEEEEE  RR        RR   CCCCCCCCCC     00000000     1111111111    GGGGGGGGGG



                    JJJJJJJJJJ   6666666666         444                                                AAAAAAAAAA
                    JJJJJJJJJJ  666666666666       4444                                               AAAAAAAAAAAA
                        JJ      66        66      44 44                                               AA        AA
                        JJ      66               44  44                                               AA        AA
                        JJ      66              44   44                                               AA        AA
                        JJ      66666666666    44444444444                                            AAAAAAAAAAAA
                        JJ      666666666666  444444444444                                            AAAAAAAAAAAA
                        JJ      66        66         44                                               AA        AA
                  JJ    JJ      66        66         44                                               AA        AA
                  JJ    JJ      66        66         44                                               AA        AA
                  JJJJJJJJ      666666666666         44                                               AA        AA
                   JJJJJJ        6666666666          44                                               AA        AA


****A   END   JOB   64  HERC01G                         ROOM        2.45.25 PM 29 JUL 23  PRINTER1  SYS TK4-  JOB   64   END   A****
****A   END   JOB   64  HERC01G                         ROOM        2.45.25 PM 29 JUL 23  PRINTER1  SYS TK4-  JOB   64   END   A****
****A   END   JOB   64  HERC01G                         ROOM        2.45.25 PM 29 JUL 23  PRINTER1  SYS TK4-  JOB   64   END   A****
****A   END   JOB   64  HERC01G                         ROOM        2.45.25 PM 29 JUL 23  PRINTER1  SYS TK4-  JOB   64   END   A****
```

## References 

