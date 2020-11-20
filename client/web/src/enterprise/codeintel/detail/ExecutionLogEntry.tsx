import React, { FunctionComponent } from 'react'
import { DateTime } from '../../../../../shared/src/graphql/schema'

export interface ExecutionLog {
    key: string
    command: string[]
    out: string
    startTime: DateTime
    exitCode: number
    durationMilliseconds: number
}

export interface ExecutionLogEntryProps {
    entry: ExecutionLog
    className?: string
}

export const ExecutionLogEntry: FunctionComponent<ExecutionLogEntryProps> = ({ entry, className }) => (
    <>
        <div className={className}>
            <p>key: {entry.key}</p>
            <p>command: {entry.command.join(' ')}</p>
            <p>startTime: {entry.startTime}</p>
            <p>exitCode: {entry.exitCode}</p>
            <p>durationMilliseconds: {entry.durationMilliseconds}</p>

            <pre className="bg-code rounded p-3">
                {/* TODO - don't do danger here */}
                <code dangerouslySetInnerHTML={{ __html: entry.out }} />
            </pre>
        </div>
    </>
)
