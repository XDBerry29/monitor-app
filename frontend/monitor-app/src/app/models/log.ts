export enum SeverityLevel {
    DEBUG = 0,
    INFO ,
    WARNING ,
    ERROR ,
    CRITICAL
  }


export interface Log {
    severity: SeverityLevel;
    time: string;
    process: string;
    message: string;
  }
