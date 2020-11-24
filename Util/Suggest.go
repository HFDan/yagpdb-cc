{{/*
    Trigger Type: Command (mention/cmd prefix)
    Trigger: "suggest"    
*/}}

{{/*This command is under the GNU GPL v3.0 License.
You may: Modify it,
Distribute it,
Patent it (provided you do not bring suit for patent infringement against anyone for making, using or distributing their own works based on the Program).

All of that as long as you: Disclose source,
Release your modification under the same GNU GPL v3.0 license
and state every modification you made.

For more info on GNU GPL v3.0 visit https://choosealicense.com/licenses/gpl-3.0/ */}}

{{/*CONFIG START*/}}
{{$channel := 780464794514882560}}                        {{/*channel to send message to*/}}
{{$memberRole := 779792628673544232}}              {{/*Member role. If you want everyone to run the command, set it to the ID of the @everyone role*/}}
{{$noMemberRoleErrorMsg := "Make a new line using \n"}} {{/*Error message to send when user does not have the Member role*/}}
{{$embedColor := 0x05B8CC}}                                     {{/*The color of the embed*/}}
{{$reactionYes := "✅"}}                                                {{/*The reaction to act as "Yes"  NOTE: If the reaction is a custom emoji, write it in this way: ":emojiName:EmojiID"*/}}
{{$reactionNo := "❌"}}                                                 {{/*The reaction to act as "No"   NOTE: If the reaction is a custom emoji, write it in this way: ":emojiName:EmojiID"*/}}
{{/*CONFIG END*/}}
 
{{/*Argument Parsing*/}}
{{$args := parseArgs 1 "Usage: -suggest <suggestion>" (carg "string" "your suggestion")}}
 
{{/*check for argc*/}}
{{if $args.IsSet 0}}
{{deleteTrigger 0}}
        {{if targetHasRoleID .User.ID 779792628673544232}}
            {{$suggestionMessage := sendMessageRetID nil (complexMessage "content" "" "embed" (cembed "title" "Sugestie:" "description" (printf "%s"  ($args.Get 0)) "timestamp" (currentTime) "author" (sdict "name" (.User.Username) "icon_url" (printf "https://cdn.discordapp.com/avatars/%s/%s.png" (toString .User.ID) (.User.Avatar))) "color" $embedColor))}}
            {{addMessageReactions nil $suggestionMessage $reactionYes $reactionNo}}
        {{else}}
            {{$errorMessage := sendMessageRetID nil $noMemberRoleErrorMsg}}
            {{deleteMessage nil $errorMessage 5}}
        {{end}}
{{end}}