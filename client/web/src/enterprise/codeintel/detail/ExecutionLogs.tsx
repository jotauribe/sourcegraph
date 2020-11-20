import React, { FunctionComponent } from 'react'
import { LsifIndexFields } from '../../../graphql-operations'
import { ExecutionLog, ExecutionLogEntry } from './ExecutionLogEntry'

export interface ExecutionLogsProps {
    index: LsifIndexFields
    className?: string
}

export const ExecutionLogs: FunctionComponent<ExecutionLogsProps> = ({ index, className }) => (
    <>
        <h3>Output logs</h3>

        {index.executionLogs.length === 0 ? (
            <>No output logs</>
        ) : (
            index.executionLogs.map((entry: ExecutionLog) => (
                <ExecutionLogEntry key={entry.key} entry={entry} className={className} />
            ))
        )}
    </>
)
